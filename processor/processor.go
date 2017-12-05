package processor

import (
	"time"

	"../gcapi"
)

// frame rate * 1000
const frameRate uint64 = 60000

var stopped bool

// ProcessAsync : Start the macro process in asynchronus
func ProcessAsync(states *[]State) {
	stopped = false
	go Process(states)
}

// Stop : Force stop the macro process
func Stop() {
	stopped = true
}

// Process : Start the macro process
func Process(states *[]State) {
	var totalFrames uint64
	start := time.Now().UnixNano()

	for _, state := range *states {
		if stopped {
			break
		}

		if &state == nil {
			continue
		}

		gcapi.Push(&state.buttons)

		totalFrames += uint64(state.frames)
		sleepTime := start + int64(totalFrames*1000*1000*1000*1000/frameRate) - time.Now().UnixNano()
		time.Sleep(time.Duration(sleepTime))
	}

	gcapi.Release()
}
