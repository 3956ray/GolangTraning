package main

import (
	"time"
)

var timer *time.Timer
var count int
var m map[string]int

func main() {
	deviceList = make([]IDevice, 1, 100)
	for {
		//	顯示菜單
		if !menu() {
			break
		}
	}

}
