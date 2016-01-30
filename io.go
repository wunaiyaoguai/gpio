package gpio

import (
	"os"
	"time"
)

type Pin struct {
	Number uint
	f      *os.File
}

func NewInput(p uint) Pin {
	pin := Pin{
		Number: p,
	}
	exportGPIO(pin)
	time.Sleep(10 * time.Millisecond)
	setDirection(pin, inDirection, 0)
	openPin(pin)
	return pin
}

func NewOutput(p uint, initHigh bool) Pin {
	pin := Pin{
		Number: p,
	}
	exportGPIO(pin)
	time.Sleep(10 * time.Millisecond)
	initVal := uint(0)
	if initHigh {
		initVal = uint(1)
	}
	setDirection(pin, outDirection, initVal)
	openPin(pin)
	return pin
}

func (p Pin) Close() {
	p.f.Close()
}