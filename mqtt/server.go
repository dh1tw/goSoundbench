package mqttserver

import (
	"log"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MqttSettings contains the Settings for the MQTT Server
type MqttSettings struct {
	Transport  string
	BrokerURL  string
	BrokerPort int
	ClientID   string
	Topics     []string
}

// MqttServer connects to an MQTT Broker and handels all incoming and outcoming
// messages
func MqttServer(s MqttSettings) {

	resultCh := make(chan ToWireMsg, 5)
	errorCh := make(chan ToWireMsg, 5)

	// handling incoming messages
	var msgHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

		task := Task{
			BaseTopic: strings.TrimSuffix(msg.Topic(), "/task"),
			Data:      msg.Payload()[:len(msg.Payload())],
			ResultCh:  resultCh,
			ErrorCh:   errorCh,
		}

		go task.Run()
	}

	// connection lost handler
	var connectionLostHandler = func(client mqtt.Client, err error) {
		log.Println("Connection lost to MQTT Broker; Reason:", err)
	}

	// since we use SetCleanSession we have to subscribe on each
	// connect or reconnect to the channels
	var onConnectHandler = func(client mqtt.Client) {
		log.Println("Connected to MQTT Broker ")

		// Subscribe to Task Topics
		for _, topic := range s.Topics {
			if token := client.Subscribe(topic+"/task", 0, nil); token.Wait() &&
				token.Error() != nil {
				log.Println(token.Error)
			}
		}
	}

	opts := mqtt.NewClientOptions().AddBroker(s.Transport + "://" + s.BrokerURL + ":" + strconv.Itoa(s.BrokerPort))
	opts.SetClientID(s.ClientID)
	opts.SetDefaultPublishHandler(msgHandler)
	opts.SetKeepAlive(time.Second * 5)
	opts.SetMaxReconnectInterval(time.Second * 2)
	opts.SetCleanSession(true)
	opts.SetOnConnectHandler(onConnectHandler)
	opts.SetConnectionLostHandler(connectionLostHandler)
	opts.SetAutoReconnect(true)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	// this loop should never quit (execpt on system exit)
	for {
		select {
		case msg := <-resultCh:
			token := client.Publish(msg.Topic, 0, false, msg.Data)
			token.Wait()
		case msg := <-errorCh:
			token := client.Publish(msg.Topic, 0, false, msg.Data)
			token.Wait()
		}
	}
}
