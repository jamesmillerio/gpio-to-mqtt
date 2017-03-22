package main

//MQTT represents the structure for our MQTT settingsin our .config file.
type MQTT struct {
	Broker string
	Port   int
}
