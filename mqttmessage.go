package main

import "encoding/json"

//MqttMessage prints the window/door
//status changes out to the terminal.
type MqttMessage struct {
	Status bool
	Name   string
}

//NewMqttMessage creates a new instance
//of our terminal receiver, which writes status changes
//out to the terminal.
func NewMqttMessage(status bool, name string) *MqttMessage {

	message := new(MqttMessage)

	message.Name = name
	message.Status = status

	return message
}

//ToBytes outputs the current message to a byte array.
func (m *MqttMessage) ToBytes() []byte {

	if message, err := json.Marshal(m); err != nil {
		panic(err)
	} else {
		return message
	}

}
