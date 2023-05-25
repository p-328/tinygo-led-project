package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	blueLed := machine.D5
	redLed := machine.D6
	orangeLed := machine.D4
	greenLed := machine.D3
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	blueLed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	greenLed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	orangeLed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	redLed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 50)
		led.High()
		time.Sleep(time.Millisecond * 50)
		led.Low()
		blueLed.Low()
		time.Sleep(time.Millisecond * 50)
		blueLed.High()
		time.Sleep(time.Millisecond * 50)
		blueLed.Low()
		orangeLed.Low()
		time.Sleep(time.Millisecond * 50)
		orangeLed.High()
		time.Sleep(time.Millisecond * 50)
		orangeLed.Low()
		greenLed.Low()
		time.Sleep(time.Millisecond * 50)
		greenLed.High()
		time.Sleep(time.Millisecond * 50)
		greenLed.Low()
		redLed.Low()
		time.Sleep(time.Millisecond * 50)
		redLed.High()
		time.Sleep(time.Millisecond * 50)
		redLed.Low()
		greenLed.Low()
		time.Sleep(time.Millisecond * 50)
		greenLed.High()
		time.Sleep(time.Millisecond * 50)
		greenLed.Low()
		orangeLed.Low()
		time.Sleep(time.Millisecond * 50)
		orangeLed.High()
		time.Sleep(time.Millisecond * 50)
		orangeLed.Low()
		blueLed.Low()
		time.Sleep(time.Millisecond * 50)
		blueLed.High()
		time.Sleep(time.Millisecond * 50)
		blueLed.Low()
		time.Sleep(time.Millisecond * 50)
	}
}
