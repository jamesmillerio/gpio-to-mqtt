package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

//SecuritySystem is the main entry point into interacting with
//our Raspberry Pi and its associated monitored windows/doors.
type SecuritySystem struct {
	configuration *Configuration
	receivers     []IStatusReceiver
	terminal      *TerminalStatusReceiver
	mqtt          *MqttStatusReceiver
}

//NewSecuritySystem initializes our security system and the Raspberry Pi.
func NewSecuritySystem(c *Configuration) *SecuritySystem {
	s := new(SecuritySystem)

	s.configuration = c
	s.receivers = []IStatusReceiver{}

	return s
}

//GetCurrentSwitchValues gets the GPIO pins with updated values so
//they can be inspected by the calling code.
func (s *SecuritySystem) GetCurrentSwitchValues() []Pin {

	return s.configuration.Pins

}

func (s *SecuritySystem) AddReceiver(receiver IStatusReceiver) {

	if receiver == nil {
		return
	}

	s.receivers = append(s.receivers, receiver)
}

//Close disposes of our Raspberry Pi's resources.
func (s *SecuritySystem) Close() {
	fmt.Print("Shutting down Raspberry Pi...\n")
	defer rpio.Close()
}

//BeginUpdating polls the Raspberry Pi for changes in values and notifies delegates as needed.
func (s *SecuritySystem) BeginUpdating() {

	go func() {

		fmt.Printf("Beginning Raspberry Pi polling...\n")

		err := rpio.Open()

		if err != nil {
			panic(err)
		}

		//Configure all of our pins.
		for i, pin := range s.configuration.Pins {
			//pin.Configure()

			s.configuration.Pins[i].Pin = rpio.Pin(pin.GPIOPin)

			fmt.Printf("Configuring pin %v...\n", s)

			//Set pin as an input pin.
			s.configuration.Pins[i].Pin.Input()

			//Set the pull of the pin.
			switch pin.Pull {
			case rpio.PullUp:
				s.configuration.Pins[i].Pin.PullUp()
			case rpio.PullDown:
				s.configuration.Pins[i].Pin.PullDown()
			case rpio.PullOff:
				s.configuration.Pins[i].Pin.PullOff()
			}

		}

		go func() {
			for {

				for i, pin := range s.configuration.Pins {

					value := pin.Pin.Read()
					prior := pin.Value
					current := value == 1

					s.configuration.Pins[i].Value = current

					if prior != current {

						//We've had a change. Notify each receiver.
						for _, receiver := range s.receivers {
							receiver.Notify(s.configuration.Pins[i])
						}
					}

				}

				time.Sleep(time.Millisecond * time.Duration(s.configuration.PollingIntervalMs))
			}
		}()

	}()
}
