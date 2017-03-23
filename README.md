# gpio-to-mqtt

<img src="./assets/raspberrypi.png" alt="Raspberry Pi" height="100"> <img src="./assets/golang.png" alt="Golang" height="100"> <img src="./assets/mqtt.png" alt="MQTT" height="100">

A simple project for translating changes to Raspberry Pi GPIO pins to MQTT messages written in Go.

This project can be handy for anything where a GPIO pin will be turned on/off and you’d like to notify listeners on your network of the change. A good example is wiring a Raspberry Pi up to door and window switches from an old alarm system to notify listeners on your network of changes. Someone opens or closes a door/window? You’ll get an MQTT message when it opens and another when it closes.

## Getting Started

This project is written in Go so you’ll need to install Go on your Raspberry Pi by running:

> sudo apt-get install golang

That should get you running with Go. Then you will need to configure your Go environment by creating your Go directories and setting your $GOPATH.

Once that’s done, you can clone the repo into your source directory and run:

> go get  
> go install

This will get the dependencies and install in your bin directory. Now you need a configuration file.

## Configuration

By default, gpio-to-mqtt will look in the directory it was launched from for a *.config* file if one was not specified as the first command line argument. A sample is provided in the root of the project titled *config.default.json.*

Options within this file include:

```json
{
  "PollingIntervalMs": 500,
  "MQTT": {
    "Broker": "127.0.0.1",
    "Port": 1883
  },
 "Pins": [{
   "GPIOPin": 22,
   "Topic": "downstairs/frontdoor",
   "Pull": 2,
   "Name": "Front Door",
   "Retain": true
 }]
}
```

You can define as many additional pins as your board supports by just adding to the pins array in the configuration file.

### Configuration Options

- **PollingIntervalMs:** The frequency with which to check for pin state changes in milliseconds.
- **MQTT**
  - **Broker:** The MQTT broker host name or ip.
  - **Port:** The MQTT broker listening port.
- **Pins**
  - **GPIOPin:** The GPIO pin number to monitor (NOT the BCM pin number).
  - **Topic:** The topic to broadcast the state change message to.
  - **Pull:** Whether to pull up (2), down (1), or off (0).
  - **Name:** A friendly name for the pin.
  - **Retain:** Whether the broker should retain state change messages.


## Extensibility

There are two main ways to extend this project for your needs. The first and most obvious way is to create your own MQTT listener (perhaps using the [yosssi/gmq](https://github.com/yosssi/gmq) library). You can just run gpio-to-mqtt as a service to publish messages and your listener can act on those messages.

However, if you'd like to extend gpio-to-mqtt for your own needs, you can do so by implementing a struct that implements IStatusReceiver. It is a very simple interface with one required method:

> type IStatusReceiver interface {  
> 	Notify(pin Pin)  
> }  

Notify is called when a pin changes states. After implementing your struct, you can register it in the main.go file by calling the .AddReceiver() method of the SecuritySystem struct. This can be done after the Terminal and MQTT status receiver are registered. That should be all there is to it.

## To Do

- [ ] Add more MQTT configuration options. We use the [yosssi/gmq](https://github.com/yosssi/gmq) library for our MQTT needs and it provides a lot of options that aren't currently supported in gpio-to-mqtt.
- [ ] Testing... This started out as a proof-of-concept so tests were foregone during that time. This needs to be rectified.
