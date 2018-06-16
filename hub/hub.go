package hub

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/go-resty/resty"
)

type (
	Hub struct {
		Options
		apiKey    string
		deviceID  string
		projectID string
		resty     *resty.Client
		client    mqtt.Client
		logger    *log.Logger
	}

	Options struct {
		MessageHandler MessageHandler
	}

	Key struct {
		Value     string `json:"value"`
		ProjectID string `json:"project_id"`
	}

	ConnectHandler func()

	MessageHandler func(topic string, message []byte)
)

func New(apiKey, deviceID string) *Hub {
	return NewWithOptions(apiKey, deviceID, Options{})
}

func NewWithOptions(apiKey, deviceID string, options Options) (h *Hub) {
	h = &Hub{
		apiKey:   apiKey,
		deviceID: deviceID,
		resty:    resty.New().SetHostURL("https://api.labstack.com").SetAuthToken(apiKey),
		logger:   log.New("hub"),
	}
	h.Options = options
	return
}

func (h *Hub) normalizeDeviceID() string {
	return fmt.Sprintf("%s:%s", h.projectID, h.deviceID)

}

func (h *Hub) normalizeTopic(name string) string {
	return fmt.Sprintf("%s/%s", h.projectID, name)
}

func (h *Hub) denormalizeTopic(name string) string {
	return strings.TrimPrefix(name, h.projectID+"/")
}

func (h *Hub) Connect() error {
	return h.ConnectWithHandler(nil)
}

func (h *Hub) ConnectWithHandler(handler ConnectHandler) error {
	key := new(Key)
	res, err := h.resty.R().
		SetResult(key).
		Get("/keys")
	if err != nil || res.StatusCode() < 200 || res.StatusCode() >= 300 {
		return errors.New("Unable to find the project")
	}
	h.projectID = key.ProjectID

	o := mqtt.NewClientOptions().
		AddBroker("tcp://hub.labstack.com:1883").
		SetUsername(h.projectID).
		SetPassword(h.apiKey).
		SetClientID(h.normalizeDeviceID())
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
