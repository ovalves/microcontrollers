package main

import (
	"machine"
	"time"
)

func NewLed(led machine.Pin) machine.Pin {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return led
}

func main() {
	var ledDelay uint32 = 10000
	green := NewLed(machine.D13)
	yellow := NewLed(machine.D12)
	red := NewLed(machine.D11)

	for {
		red.Set(true)
		delay(ledDelay)

		yellow.Set(true)
		delay(2000)

		green.Set(true)
		red.Set(false)
		yellow.Set(false)
		delay(ledDelay)

		yellow.Set(true)
		green.Set(false)
		delay(2000)

		yellow.Set(false)
	}

}

func delay(t uint32) {
	time.Sleep(time.Duration(1000000 * t))
}
