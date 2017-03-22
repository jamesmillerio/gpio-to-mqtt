package main

//IStatusReceiver defines an interface for receiving
//notifications when a door or window opens/closes.
type IStatusReceiver interface {
	Notify(pin Pin)
}
