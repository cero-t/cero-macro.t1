package processor

// State : state for PS controller
type State struct {
	buttons [36]int8
	frames  uint16
}
