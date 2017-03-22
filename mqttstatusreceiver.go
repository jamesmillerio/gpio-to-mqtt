package main

import (
	"fmt"
	"time"

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
			break
		}
	}
}

//Notify prints the status change to the terminal.
func (m *MqttStatusReceiver) Notify(pin Pin) {

	fmt.Printf("MQTT Receiver: %v.\n", pin)

}

/*for {

	pin := <-notifications

	message := "ON"

	if pin.Value {
		message = "OFF"
	}

	topic := pin.Topic

	if len(topic) <= 0 {
		topic = strings.ToLower("home-assistant/window/contact/" + strings.Replace(pin.Name, " ", "-", -1))
	}

	fmt.Printf("Topic: %v\n", topic)

	err := mqttClient.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte(topic),
		Message:   []byte(message),
	})

	if err != nil {
		panic(err)
	}

}*/
