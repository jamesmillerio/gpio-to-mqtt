package main

import "fmt"

//TerminalStatusReceiver prints the window/door
//status changes out to the terminal.
type TerminalStatusReceiver struct {
	configuration *Configuration
	notifications chan Pin
}

//NewTerminalStatusReceiver creates a new instance
//of our terminal receiver, which writes status changes
//out to the terminal.
func NewTerminalStatusReceiver(configuration *Configuration, notificationsChannel chan Pin) *TerminalStatusReceiver {

	receiver := new(TerminalStatusReceiver)

	receiver.configuration = configuration
	receiver.notifications = notificationsChannel

	receiver.listen()

	return receiver
}

//listen begins listening for channel events.
func (t *TerminalStatusReceiver) listen() {
	go func() {

		for {

			t.Notify(<-t.notifications)

		}

	}()
}

//Notify prints the status change to the terminal.
func (t *TerminalStatusReceiver) Notify(pin Pin) {

	fmt.Printf("Terminal Receiver: %v.\n", pin)

}
