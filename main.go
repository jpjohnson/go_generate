package main

import (
	"fmt"

	"github.com/jpjohnson/go_generate/painkiller"
)

func main() {
	pill1 := painkiller.Ibuprofen
	pill2 := painkiller.Acetaminohpen
	fmt.Println("pill 1 :", pill1.String())
	fmt.Println("pill 2 :", pill2.String())
}
