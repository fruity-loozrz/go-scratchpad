package main

import (
	"fmt"
	"log"

	"github.com/fruity-loozrz/go-scratchpad/internal/vnljs"
)

func main() {
	api, err := vnljs.ExecuteVnlJs(`
	api
		.Action()
		.Dur(1)
		.Easing("sharp")
		.Platter(0, 1/2)
		.FaderPattern("open")`)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Printf("API: %+v\n", api)
	for i, action := range api.Actions() {
		fmt.Printf("Action %d: %+v\n", i, action)
	}
}
