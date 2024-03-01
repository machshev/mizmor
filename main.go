package main

import (
	"fmt"

	"github.com/machshev/mizmor/generate"
	"github.com/machshev/mizmor/play"
	"github.com/machshev/mizmor/scale"
)

func main() {
	var err error
	var gen = generate.NewGenerator(scale.NewPsalmScaleSHV())

	err = gen.GenMidi("out")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = play.Play("out")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}
