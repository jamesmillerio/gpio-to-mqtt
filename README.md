# gpio-to-mqtt

![Golang](./assets/golang.png  =100x100) ![Raspberry Pi](./assets/raspberrypi.png =100x100) ![MQTT](./assets/mqtt.png  =100x100)

A simple project for translating changes to Raspberry Pi GPIO pins to MQTT messages written in Go.

This project can be handy for anything where a GPIO pin will be turned on/off and you’d like to notify listeners on your network of the change. A good example is wiring a Raspberry Pi up to door and window switches from an old alarm system to notify listeners on your network of changes. Someone opens or closes a window? You’ll get an MQTT message.

## Getting Started
---
This project is written in Go so you’ll need to install go on your Raspberry Pi by running:

> sudo apt-get install golang

That should get your running with Go. Then you will need to configure your Go environment by creating your Go directories and setting your $GOPATH.

Once that’s done, you can clone the repo into your source directory and run:

> go get
> go install

This will get the dependencies and install in your bin directory.

Before you can run, however, you’ll need to create a configuration file. By default, gpio-to-mqtt will look in the directory it was launched from for a *.config* file if one was not specified. You can specify your own by passing it in as the first argument when executing the application. A sample is provided in the root of the project titled *config.default.json.*

Options within this file include:

> {
>   "PollingIntervalMs": 500, 	      //How often (in ms) to poll the GPIO pins for state changes.
>   "MQTT": {
>     "Broker": "127.0.0.1”,		      //The MQTT broker host address.
>     “Port”: 1883, 				          //The MQTT broker host port.
>  },
>  "Pins": [{ 					              //An array of GPIO pins to monitor.
>    "GPIOPin": 22,			              //The GPIO pin number, NOT the BCM pin.
>    "Topic": "downstairs/frontdoor",	//The topic to publish to on a state change.
>    "Pull": 2,					              //Whether to pull up (2), down (1), or off (0).
>    "Name": "Front Door",		        //A friendly name for the pin.
>    "Retain": true				            //Whether the MQTT broker should retain the messages.
>  }]
> }


You can define as many additional pins as your board supports by just adding to the pins array in the configuration file.
