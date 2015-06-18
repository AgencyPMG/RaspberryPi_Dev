// A simple program that blinks the Green LED on the Pi when touching GPIO22
// and exits when GPIO20 is touched.

package main

import (
	"fmt"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"time"
)

func main() {
	if err := embd.InitGPIO(); err != nil {
		panic(err)
	}
	defer embd.CloseGPIO()

	pin, _ := embd.NewDigitalPin(20)
	pin2, _ := embd.NewDigitalPin(21)
	defer pin.Close()
	defer pin2.Close()

	pin.SetDirection(embd.In)
	pin2.SetDirection(embd.In)
	fmt.Println("**********************************")
	fmt.Println("**     Listening on GPIO22      **")
	fmt.Println("** Do not CTRLC out! Use GPIO20 **")
	fmt.Println("**********************************")
	for {
		if isPinHigh(pin) == true {
			break
		}
		if isPinHigh(pin2) == true {
			embd.LEDToggle("LED0")
			time.Sleep(100 * time.Millisecond)
		} else {
			embd.LEDOff("LED0")
		}
	}
}

func isPinHigh(pin embd.DigitalPin) bool {
	val, err := pin.Read()
	if err != nil {
		return false
	}
	if val == 1 {
		return true
	}
	return false
}
