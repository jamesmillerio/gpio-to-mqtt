package main

import "fmt"

//TerminalStatusReceiver prints the window/door
//status changes out to the terminal.
type TerminalStatusReceiver struct {
	configuration *Configuration
}

//NewTerminalStatusReceiver creates a new instance
//of our terminal receiver, which writes status changes
//out to the terminal.
func NewTerminalStatusReceiver(configuration *Configuration) *TerminalStatusReceiver {

	receiver := new(TerminalStatusReceiver)

	receiver.configuration = configuration

	return receiver
}

//Notify prints the status change to the terminal.
func (t *TerminalStatusReceiver) Notify(pin Pin) {

	fmt.Printf("Terminal Receiver: Pin: %v Value: %v\n", pin.GPIOPin, pin.Value)

}
