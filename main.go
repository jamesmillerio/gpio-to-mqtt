package main

import (
	"math"
	"os"
	"os/signal"
	"sync"

	"github.com/sasbury/logging"
)

var config *Configuration
var securitySystem *SecuritySystem
var waitGroup sync.WaitGroup

func main() {

	//COmnfigure our wait group
	waitGroup.Add(1)

	//Set up some logging
	rollingFile := logging.NewRollingFileAppender("events", "log", math.MaxInt64, 1)

	config = NewSecurityConfiguration()

	logging.AddAppender(rollingFile)

	securitySystem = NewSecuritySystem(config)

	//If the app closes by normal means, shut everything down.
	defer securitySystem.Close()

	handleInterrupts()

	//notifications := securitySystem.BeginUpdating()
	securitySystem.BeginUpdating()

	waitGroup.Wait()

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
