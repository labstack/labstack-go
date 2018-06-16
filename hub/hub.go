package hub

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/go-resty/resty"
)

type (
	Hub struct {
		Options
		resty  *resty.Client
		key    *Key
		client mqtt.Client
		logger *log.Logger
	}

	Options struct {
		DeviceID       string
		MessageHandler MessageHandler
	}

	Key struct {
		ID        string `json:"id"`
		ProjectID string `json:"project_id"`
	}

	ConnectHandler func()

	MessageHandler func(topic string, message []byte)
)

func New(apiKey string) *Hub {
	return NewWithOptions(apiKey, Options{})
}

func NewWithOptions(apiKey string, options Options) (h *Hub) {
	h = &Hub{
		key: &Key{
			ID: apiKey,
		},
		resty:  resty.New().SetHostURL("https://api.labstack.com").SetAuthToken(apiKey),
		logger: log.New("hub"),
	}
	h.Options = options
	if h.DeviceID == "" {
		h.DeviceID, _ = os.Hostname()
	}
	res, err := h.resty.R().
		SetResult(h.key).
		Get("/keys/" + h.key.ID)
	if err != nil {
		h.logger.Fatal(err)
	}
	if res.StatusCode() < 200 || res.StatusCode() >= 300 {
		h.logger.Fatal(err)
	}
	h.DeviceID = h.normalizeDeviceID(h.DeviceID)
	return
}

func (h *Hub) normalizeDeviceID(id string) string {
	return fmt.Sprintf("%s:%s", h.key.ProjectID, id)

}

func (h *Hub) normalizeTopic(name string) string {
	return fmt.Sprintf("%s/%s", h.key.ProjectID, name)
}

func (h *Hub) denormalizeTopic(name string) string {
	return strings.TrimPrefix(name, h.key.ProjectID+"/")
}

func (h *Hub) Connect() error {
	return h.ConnectWithHandler(nil)
}

func (h *Hub) ConnectWithHandler(handler ConnectHandler) error {
	o := mqtt.NewClientOptions().
		AddBroker("tcp://hub.labstack.com:1883").
		SetUsername(h.key.ProjectID).
		SetPassword(h.key.ID).
		SetClientID(h.DeviceID)
	if handler != nil {
		o.OnConnect = func(_ mqtt.Client) {
			handler()
		}
	}
	h.client = mqtt.NewClient(o)
	t := h.client.Connect()
	t.Wait()
	return t.Error()
}

func (h *Hub) Publish(topic string, message interface{}) error {
	t := h.client.Publish(h.normalizeTopic(topic), 0, false, message)
	t.Wait()
	return t.Error()
}

func (h *Hub) Subscribe(topic string) error {
	return h.SubscribeWithHandler(topic, nil)
}

func (h *Hub) SubscribeWithHandler(topic string, handler MessageHandler) error {
	t := h.client.Subscribe(h.normalizeTopic(topic), 0, func(_ mqtt.Client, m mqtt.Message) {
		topic := h.denormalizeTopic(m.Topic())
		if handler != nil {
			handler(topic, m.Payload())
		}
		if h.MessageHandler != nil {
			h.MessageHandler(topic, m.Payload())
		}
	})
	t.Wait()
	return t.Error()
}

func (h *Hub) Unsubscribe(topic string) error {
	t := h.client.Unsubscribe(h.normalizeTopic(topic))
	t.Wait()
	return t.Error()
}

func (h *Hub) Disconnect() {
	h.client.Disconnect(1000)
}

func (h *Hub) Run() {
	for {
		time.Sleep(time.Second)
	}
}
