package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

//MqttStatusReceiver prints the window/door
//status changes out to the terminal.
type MqttStatusReceiver struct {
	configuration *Configuration
	mqttClient    *client.Client
}

//NewMqttStatusReceiver creates a new instance
//of our terminal receiver, which writes status changes
//out to the terminal.
func NewMqttStatusReceiver(configuration *Configuration) *MqttStatusReceiver {

	receiver := new(MqttStatusReceiver)

	receiver.configuration = configuration

	receiver.connectToMqttBroker()

	return receiver
}

func (m *MqttStatusReceiver) connectToMqttBroker() {

	// Create an MQTT Client.
	m.mqttClient = client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			fmt.Println(err)
		},
	})

	// Terminate the Client.
	defer m.mqttClient.Terminate()

	retries := 5

	for attempt := 0; attempt < retries; attempt++ {
		// Connect to the MQTT Server.
		err := m.mqttClient.Connect(&client.ConnectOptions{
			Network:  "tcp",
			Address:  fmt.Sprintf("%v:%v", m.configuration.MQTT.Broker, m.configuration.MQTT.Port),
			ClientID: []byte("rpi-security-system"),
		})

		if err != nil {
			fmt.Print("An error occurred while trying to connect to our MQTT server. Retrying in 10 seconds...")
			time.Sleep(time.Millisecond * time.Duration(10000))
		} else {

			//Now that we're connected, send some required messages.
			for _, pin := range m.configuration.Pins {

				var nameTopic = m.appendTopic(pin.Topic, "name")
				var statusTopic = m.appendTopic(pin.Topic, "status")
				var deviceClassTopic = m.appendTopic(pin.Topic, "deviceClass")
				var identifierTopic = m.appendTopic(pin.Topic, "identifier")
				var status = "0"

				if pin.Value {
					status = "1"
				}

				m.sendMessageString(nameTopic, pin.Name, true)
				m.sendMessageString(statusTopic, status, true)
				m.sendMessageString(deviceClassTopic, pin.DeviceClass, true)
				m.sendMessageString(identifierTopic, pin.Identifier, true)
			}

			break
		}
	}
}

//Notify prints the status change to the terminal.
func (m *MqttStatusReceiver) Notify(pin Pin) {

	//message := NewMqttMessage(pin.Value, pin.Name)
	var statusTopic = m.appendTopic(pin.Topic, "status")
	var status = "0"

	if pin.Value {
		status = "1"
	}

	//m.sendMessageBytes(pin.Topic, message.ToBytes(), pin.Retain)
	m.sendMessageString(statusTopic, status, pin.Retain)

}

//Notify prints the status change to the terminal.
func (m *MqttStatusReceiver) sendMessageString(topic string, message string, retain bool) {

	m.sendMessageBytes(topic, []byte(message), retain)

}

func (m *MqttStatusReceiver) sendMessageBytes(topic string, message []byte, retain bool) {

	err := m.mqttClient.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte(topic),
		Message:   message,
		Retain:    retain,
	})

	if err != nil {
		panic(err)
	}
}

func (m *MqttStatusReceiver) appendTopic(topic string, child string) string {

	var separator = "/"
	var topics = strings.Split(topic, separator)

	topics = append(topics, child)

	return strings.Join(topics, separator)
}
