package abr

import (
	"fmt"
	"math/rand"
	"time"
)

type House struct {
	window_closed bool
	fan_off       bool
}

func (h *House) SetAlarm() {
	fmt.Println("Alarme armée")
	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("L'alarme a terminé son compte à rebours")
		return
	}()
}

func (h *House) StartEnvironnement(c chan (string)) {

	for order := range c {
		switch order {
		case "window":
			if !h.window_closed {
				time.Sleep(time.Duration(10 + rand.Intn(31)))
				h.window_closed = true
				c <- "Window is closed"
				//fmt.Println("window closed")
			} else {
				c <- "Window was already closed"
				//fmt.Println("window already closed")
			}
		case "fan":
			if !h.fan_off {
				time.Sleep(time.Duration(10 + rand.Intn(31)))
				h.fan_off = true
				c <- "Fan is off"
				//fmt.Println("Fan off")

			} else {
				c <- "Fan was already off"
				//fmt.Println("Fan was already off")

			}
		}
	}
}
