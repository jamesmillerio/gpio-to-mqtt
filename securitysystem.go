package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

//SecuritySystem is the main entry point into interacting with
//our Raspberry Pi and its associated monitored windows/doors.
type SecuritySystem struct {
	configuration       *Configuration
	receivers           []IStatusReceiver
	terminal            *TerminalStatusReceiver
	mqtt                *MqttStatusReceiver
	notificationChannel chan Pin
}

//NewSecuritySystem initializes our security system and the Raspberry Pi.
func NewSecuritySystem(c *Configuration) *SecuritySystem {
	s := new(SecuritySystem)

	s.configuration = c
	s.notificationChannel = make(chan Pin)
	s.terminal = NewTerminalStatusReceiver(c, s.notificationChannel)
	s.mqtt = NewMqttStatusReceiver(c, s.notificationChannel)
	s.receivers = []IStatusReceiver{s.terminal}

	return s
}

//GetCurrentSwitchValues gets the GPIO pins with updated values so
//they can be inspected by the calling code.
func (s *SecuritySystem) GetCurrentSwitchValues() []Pin {

	return s.configuration.Pins

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
			if pin.Pull {
				s.configuration.Pins[i].Pin.PullUp()
			} else {
				s.configuration.Pins[i].Pin.PullDown()
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
						s.notificationChannel <- s.configuration.Pins[i]
					}

				}

				time.Sleep(time.Millisecond * time.Duration(s.configuration.PollingIntervalMs))
			}
		}()

	}()
}
