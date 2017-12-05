package gcapi

import (
	"errors"
	"log"
	"time"
)

// Init : Initialize the device
func Init() error {
	loadDll()

	connected := false
	for i := 0; i < 3; i++ {
		connected = isConnected()
		if connected {
			break
		}
		log.Println("Detecting Titan One...")
		time.Sleep(100 * time.Millisecond)
	}

	if !connected {
		return errors.New("Titan One is not connected")
	}

	ver := getFWVer()
	log.Println("FW Version:", ver)

	return nil
}

// Push : push buttons
func Push(inputs *[36]int8) bool {
	return write(inputs)
}

// Release : release all buttons
func Release() bool {
	var inputs [36]int8
	return write(&inputs)
}

// Close : close connection to the device
func Close() {
	unload()
}
