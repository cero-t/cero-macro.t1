package processor

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	blanks    = regexp.MustCompile(`\s+`)
	nums      = regexp.MustCompile(`^\d+$`)
	loopStart = regexp.MustCompile(`<loop\s(\d+)\s*>`)
)

const (
	buttonPS       = 0
	buttonSHARE    = 1
	buttonSELECT   = 1
	buttonOPTIONS  = 2
	buttonSTART    = 2
	buttonR1       = 3
	buttonR2       = 4
	buttonR3       = 5
	buttonL1       = 6
	buttonL2       = 7
	buttonL3       = 8
	buttonRX       = 9
	buttonRY       = 10
	buttonLX       = 11
	buttonLY       = 12
	buttonUp       = 13
	buttonDown     = 14
	buttonLeft     = 15
	buttonRight    = 16
	buttonTriangle = 17
	buttonCircle   = 18
	buttonCross    = 19
	buttonSquare   = 20
	buttonACCX     = 21
	buttonACCY     = 22
	buttonACCZ     = 23
	buttonGYROX    = 24
	buttonGYROY    = 25
	buttonGYROZ    = 26
	buttonTOUCH    = 27
	buttonTOUCHX   = 28
	buttonTOUCHY   = 29
)

func TextToStates(text *string) *[]State {
	body := strings.Replace(*text, "\r\n", "\n", -1)
	body = strings.Replace(body, "\r", "\n", -1)
	lines := strings.Split(body, "\n")
	return LinesToStates(&lines)
}

func LinesToStates(lines *[]string) *[]State {
	states := make([]State, 0, len(*lines))
	buffer := make([]State, 0, len(*lines))
	var count int

	for _, v := range *lines {
		if loopStart.MatchString(v) {
			count, _ = strconv.Atoi(loopStart.FindStringSubmatch(v)[1])
			states = append(states, buffer...)
			buffer = make([]State, 0, len(*lines))
			continue
		} else if strings.Index(v, "</loop>") == 0 {
			for i := 0; i < count; i++ {
				states = append(states, buffer...)
			}
			buffer = make([]State, 0, len(*lines))
			continue
		}

		state, err := lineToState(&v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if state == nil {
			continue
		}

		buffer = append(buffer, *state)
	}

	states = append(states, buffer...)

	return &states
}

func lineToState(line *string) (*State, error) {
	if line == nil || len(*line) == 0 {
		return nil, nil
	}

	pair := blanks.Split(*line, 2)
	if len(pair) != 2 {
		return nil, errors.New("line does not contain any blanks: " + *line)
	}

	keys := strings.Split(pair[0], ",")
	frames, err := strconv.Atoi(pair[1])
	if err != nil {
		return nil, errors.New("second value must be integer: " + *line)
	}

	state := State{}
	state.frames = uint16(frames)

	for _, v := range keys {
		if nums.MatchString(v) {
			direction, _ := strconv.Atoi(v)
			switch direction {
			case 1:
				state.buttons[buttonDown] = 100
				state.buttons[buttonLeft] = 100
			case 2:
				state.buttons[buttonDown] = 100
			case 3:
				state.buttons[buttonDown] = 100
				state.buttons[buttonRight] = 100
			case 4:
				state.buttons[buttonLeft] = 100
			case 6:
				state.buttons[buttonRight] = 100
			case 7:
				state.buttons[buttonUp] = 100
				state.buttons[buttonLeft] = 100
			case 8:
				state.buttons[buttonUp] = 100
			case 9:
				state.buttons[buttonUp] = 100
				state.buttons[buttonRight] = 100
			}
		} else if v == "lp" {
			state.buttons[buttonSquare] = 100
		} else if v == "mp" {
			state.buttons[buttonTriangle] = 100
		} else if v == "hp" {
			state.buttons[buttonR1] = 100
		} else if v == "lk" {
			state.buttons[buttonCross] = 100
		} else if v == "mk" {
			state.buttons[buttonCircle] = 100
		} else if v == "hk" {
			state.buttons[buttonR2] = 100
		} else if v == "pause" || v == "start" || v == "ooptions" {
			state.buttons[buttonSTART] = 100
		} else if v == "save" {
			state.buttons[buttonR3] = 100
		} else if v == "reload" || v == "select" || v == "share" {
			state.buttons[buttonSELECT] = 100
		}
	}

	return &state, nil
}
