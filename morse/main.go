package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.D13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		dotFor(led)
		delay(100)
		dashFor(led)
		delay(100)
		dotFor(led)
		delay(1000)
	}
}

func dotFor(led machine.Pin) {
	for i := 0; i < 3; i++ {
		led.Set(true)
		delay(150)
		led.Set(false)
		delay(100)
	}
}

func dashFor(led machine.Pin) {
	for i := 0; i < 3; i++ {
		led.Set(true)
		delay(400)
		led.Set(false)
		delay(100)
	}
}

func delay(t uint32) {
	time.Sleep(time.Duration(1000000 * t))
}
