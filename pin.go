package main

import "github.com/stianeikeland/go-rpio"

//Pin represents a pin that is used on the Raspberry Pi for GPIO.
type Pin struct {
	GPIOPin     int
	Pin         rpio.Pin
	Topic       string
	Pull        rpio.Pull
	Status      bool
	Value       bool
	Retain      bool
	Name        string
	DeviceClass string
	Identifier  string
}

//Configure loads the specified JSON file into the provided instance.
/*func (s *Pin) Configure() {

	s.Pin = rpio.Pin(s.GPIOPin)

	fmt.Printf("Configuring pin %v...\n", s)

	//Set pin as an input pin.
	s.Pin.Input()

	//Set the pull of the pin.
	if s.Pull {
		s.Pin.PullUp()
	} else {
		s.Pin.PullDown()
	}

}*/
