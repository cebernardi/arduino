package main

import (
	"machine"

	"github.com/cebernardi/arduino/pkg/dht11"
)

func main() {
	pin := machine.Pin(7)
	sens := dht11.NewSensor(pin)

	sens.Read()

}
