// +build !windows

package gcapi

import (
	"log"
	"time"
)

var (
	timer = time.Now().UnixNano()
)

func loadDll() {
	log.Println("loadDll")
}

func unload() bool {
	log.Println("unload")
	return true
}

func isConnected() bool {
	log.Println("isConnected")
	return true
}

func getFWVer() uint {
	log.Println("getFWVer")
	return 0
}

func read() bool {
	log.Println("read")
	return true
}

func write(inputs [36]int8) bool {
	log.Println("write", inputs)
	return true
}

func getTimeVal() uint {
	log.Println("getTimeVal")
	return 0
}

func calcPressTime() uint {
	log.Println("calcPressTime")
	return 0
}
