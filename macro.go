package main

import (
	"./file"
	"./gcapi"
	"./http"
	"./processor"

	"flag"
	"log"
	"net/url"
)

func main() {
	err := gcapi.Init()
	if err != nil {
		log.Fatal("gcapi is not enabled.", err)
	} else {
		log.Println("gcapi is enabled.")
	}
	defer gcapi.Close()

	flag.Parse()
	if flag.NArg() == 0 {
		serverMode()
	} else {
		textMode(flag.Arg(0))
	}
}

func textMode(path string) {
	lines := file.ReadAll(path)
	states := processor.LinesToStates(lines)
	processor.Process(states)
}

func serverMode() {
	http.Server(serverFunction)
}

func serverFunction(form *url.Values) string {
	if len(form.Get("stop")) > 0 {
		processor.Stop()
		return "stopped"
	}

	player1 := form.Get("player1")

	states := processor.TextToStates(&player1)
	processor.ProcessAsync(states)

	return "processed"
}
