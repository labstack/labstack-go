package labstack

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/shirou/gopsutil/process"
)

type (
	// Cube defines the LabStack cube service.
	Cube struct {
		process        *process.Process
		client         *Client
		requests       []*Request
		activeRequests int64
		started        int64
		mutex          sync.RWMutex

		// LabStack API key
		APIKey string

		// Node id
		Node string

		// Tags
		Tags []string

		// Number of requests in a batch
		BatchSize int

		// Interval in seconds to dispatch the batch
		DispatchInterval time.Duration

		// TODO: To be implemented
		ClientLookup string
	}

	// Request defines a request payload to be corded.
	Request struct {
		time          time.Time
		ID            string   `json:"id"`
		Time          int64    `json:"time"`
		Tags          []string `json:"tags,omitempty"`
		Host          string   `json:"host"`
		Path          string   `json:"path"`
		Method        string   `json:"method"`
		Status        int      `json:"status"`
		BytesIn       int64    `json:"bytes_in"`
		BytesOut      int64    `json:"bytes_out"`
		Latency       int64    `json:"latency"`
		ClientID      string   `json:"client_id"`
		RemoteIP      string   `json:"remote_ip"`
		UserAgent     string   `json:"user_agent"`
		Active        int64    `json:"active"`
		Error         string   `json:"error"`
		StackTrace    string   `json:"stack_trace"`
		Node          string   `json:"node"`
		Uptime        int64    `json:"uptime"`
		CPUPercent    float32  `json:"cpu_percent"`
		MemoryPercent float32  `json:"memory_percent"`
	}
)

func (c *Cube) resetRequests() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.requests = make([]*Request, 0, c.BatchSize)
}

func (c *Cube) appendRequest(r *Request) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.requests = append(c.requests, r)
}

func (c *Cube) listRequests() []*Request {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	requests := make([]*Request, len(c.requests))
	for i, r := range c.requests {
		requests[i] = r
	}
	return requests
}

func (c *Cube) requestsLength() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.requests)
}

// Dispatch dispatches the requests batch.
func (c *Cube) Dispatch() error {
	if len(c.requests) == 0 {
		return nil
	}

	err := new(APIError)
	r, e := c.client.resty.R().
		SetBody(c.listRequests()).
		SetError(err).
		Post("/cube")
	if e != nil {
		return &APIError{
			Message: e.Error(),
		}
	}

	if c.client.error(r) {
		return err
	}
	c.resetRequests()
	return nil
}

// Start starts cording an HTTP request.
func (c *Cube) Start(r *http.Request, w http.ResponseWriter) (req *Request) {
	atomic.AddInt64(&c.activeRequests, 1)

	// Daemon
	if c.started == 0 {
		go func() {
			d := time.Duration(c.DispatchInterval) * time.Second
			for range time.Tick(d) {
				if err := c.Dispatch(); err != nil {
					c.client.logger.Error(err)
				}
			}
		}()
		atomic.AddInt64(&c.started, 1)
	}

	start, _ := c.process.CreateTime()
	cpu, _ := c.process.CPUPercent()
	mem, _ := c.process.MemoryPercent()
	now := time.Now()
	req = &Request{
		time:          now,
		ID:            RequestID(r, w),
		Time:          now.UnixNano() / 1000000,
		Tags:          c.Tags,
		Host:          r.Host,
		Path:          r.URL.Path,
		Method:        r.Method,
		UserAgent:     r.UserAgent(),
		RemoteIP:      RealIP(r),
		Node:          c.Node,
		Uptime:        time.Now().Unix() - start/1000,
		CPUPercent:    float32(cpu),
		MemoryPercent: mem,
	}
	req.ClientID = req.RemoteIP
	req.Active = c.activeRequests
	cl := r.Header.Get("Content-Length")
	if cl == "" {
		cl = "0"
	}
	l, err := strconv.ParseInt(cl, 10, 64)
	if err != nil {
		c.client.logger.Error(err)
	}
	req.BytesIn = l
	c.appendRequest(req)
	return
}

// Stop stops cording an HTTP request.
func (a *Cube) Stop(r *Request, status int, size int64) {
	atomic.AddInt64(&a.activeRequests, -1)
	r.Status = status
	r.BytesOut = size
	r.Latency = int64(time.Now().Sub(r.time))

	// Dispatch batch
	if a.requestsLength() >= a.BatchSize {
		go func() {
			if err := a.Dispatch(); err != nil {
				a.client.logger.Error(err)
			}
		}()
	}
}

// RequestID returns the request ID from the request or response.
func RequestID(r *http.Request, w http.ResponseWriter) string {
	id := r.Header.Get("X-Request-ID")
	if id == "" {
		id = w.Header().Get("X-Request-ID")
	}
	return id
}

// RealIP returns the real IP from the request.
func RealIP(r *http.Request) string {
	ra := r.RemoteAddr
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		ra = strings.Split(ip, ", ")[0]
	} else if ip := r.Header.Get("X-Real-IP"); ip != "" {
		ra = ip
	} else {
		ra, _, _ = net.SplitHostPort(ra)
	}
	return ra
}
