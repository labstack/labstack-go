package hub

import (
	"fmt"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type (
	Hub struct {
		Options
		accountID string
		apiKey    string
		deviceID  string
		client    mqtt.Client
	}

	Options struct {
		MessageHandler MessageHandler
	}

	ConnectHandler func()

	MessageHandler func(topic string, message []byte)
)

func New(accountID, apiKey, deviceID string) *Hub {
	return NewWithOptions(accountID, apiKey, deviceID, Options{})
}

func NewWithOptions(accountID, apiKey, deviceID string, options Options) *Hub {
	h := &Hub{
		accountID: accountID,
		apiKey:    apiKey,
	}
	h.deviceID = h.normalizeDeviceID(deviceID)
	h.Options = options
	return h
}

func (h *Hub) normalizeDeviceID(id string) string {
	return fmt.Sprintf("%s:%s", h.accountID, id)
}

func (h *Hub) normalizeTopic(name string) string {
	return fmt.Sprintf("%s/%s", h.accountID, name)
}

func (h *Hub) denormalizeTopic(name string) string {
	return strings.TrimPrefix(name, h.accountID+"/")
}

func (h *Hub) Connect() error {
	return h.ConnectWithHandler(nil)
}

func (h *Hub) ConnectWithHandler(handler ConnectHandler) error {
	o := mqtt.NewClientOptions().
		AddBroker("tcp://hub.labstack.com:1883").
		SetUsername(h.accountID).
		SetPassword(h.apiKey).
		SetClientID(h.deviceID)
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
