package axis

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
	Axis struct {
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

func New(apiKey, deviceID string) *Axis {
	return NewWithOptions(apiKey, deviceID, Options{})
}

func NewWithOptions(apiKey, deviceID string, options Options) (a *Axis) {
	a = &Axis{
		apiKey:   apiKey,
		deviceID: deviceID,
		resty:    resty.New().SetHostURL("https://api.labstack.com").SetAuthToken(apiKey),
		logger:   log.New("axis"),
	}
	a.Options = options
	return
}

func (a *Axis) normalizeDeviceID() string {
	return fmt.Sprintf("%s:%s", a.projectID, a.deviceID)

}

func (a *Axis) normalizeTopic(name string) string {
	return fmt.Sprintf("%s/%s", a.projectID, name)
}

func (a *Axis) denormalizeTopic(name string) string {
	return strings.TrimPrefix(name, a.projectID+"/")
}

func (a *Axis) Connect() error {
	return a.ConnectWithHandler(nil)
}

func (a *Axis) ConnectWithHandler(handler ConnectHandler) error {
	// Find project id
	key := new(Key)
	res, err := a.resty.R().
		SetResult(key).
		Get("/axis/key")
	if err != nil || res.StatusCode() < 200 || res.StatusCode() >= 300 {
		return errors.New("Unable to find the project")
	}
	a.projectID = key.ProjectID

	// Connect
	o := mqtt.NewClientOptions().
		AddBroker("tcp://axis.labstack.com:1883").
		SetUsername(a.projectID).
		SetPassword(a.apiKey).
		SetClientID(a.normalizeDeviceID())
	if handler != nil {
		o.OnConnect = func(_ mqtt.Client) {
			handler()
		}
	}
	a.client = mqtt.NewClient(o)
	t := a.client.Connect()
	t.Wait()
	return t.Error()
}

func (a *Axis) Publish(topic string, message interface{}) error {
	t := a.client.Publish(a.normalizeTopic(topic), 0, false, message)
	t.Wait()
	return t.Error()
}

func (a *Axis) Subscribe(topic string) error {
	return a.SubscribeWithHandler(topic, nil)
}

func (a *Axis) SubscribeWithHandler(topic string, handler MessageHandler) error {
	t := a.client.Subscribe(a.normalizeTopic(topic), 0, func(_ mqtt.Client, m mqtt.Message) {
		topic := a.denormalizeTopic(m.Topic())
		if handler != nil {
			handler(topic, m.Payload())
		}
		if a.MessageHandler != nil {
			a.MessageHandler(topic, m.Payload())
		}
	})
	t.Wait()
	return t.Error()
}

func (a *Axis) Unsubscribe(topic string) error {
	t := a.client.Unsubscribe(a.normalizeTopic(topic))
	t.Wait()
	return t.Error()
}

func (a *Axis) Disconnect() {
	a.client.Disconnect(1000)
}

func (a *Axis) Run() {
	for {
		time.Sleep(time.Second)
	}
}
