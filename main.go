package main

import (
	"math"
	"os"
	"os/signal"

	"github.com/sasbury/logging"
)

var config *Configuration
var securitySystem *SecuritySystem

func main() {

	//Set up some logging
	rollingFile := logging.NewRollingFileAppender("events", "log", math.MaxInt64, 1)

	config = NewSecurityConfiguration()

	logging.AddAppender(rollingFile)

	securitySystem = NewSecuritySystem(config)

	//If the app closes by normal means, shut everything down.
	defer securitySystem.Close()

	//notifications := securitySystem.BeginUpdating()
	securitySystem.BeginUpdating()

	handleInterrupts()

}

//handleInterrupts handles any signal interrupts that are called
//to make sure that our Raspberry Pi gets shut down appropriately.
func handleInterrupts() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				securitySystem.Close()
				os.Exit(1)
			}

		}
	}()
}
