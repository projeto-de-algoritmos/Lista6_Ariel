package main

import (
	//"./layer1_alignmen-sequence"
	//"./native_parents"
	"./native_tape"
	//"./native_tapedist"
	"fmt"
	"os"
)

type DNA struct {
	allbuilded bool
	verified   map[string]bool
	builded    map[string]bool
	//guess_parents n_parents.parents
	children n_tape.Tape
	father   n_tape.Tape
	//distances     n_tapedist.tapedist
}

// func (t *tape) readlines(path string) error {
func (dna *DNA) setupfathers(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	} else {
		defer file.Close()
		err := dna.father.Readlines(path)
		return err
	}
}

func (dna *DNA) setupchildren(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	} else {
		defer file.Close()
		err := dna.children.Readlines(path)
		return err
	}
}
func main() {
	var dna DNA
	dna.setupfathers("Fathers.txt")
	dna.setupchildren("Children.txt")
	for i := 0; i < 10; i++ {
		fmt.Println(dna.father.Memory[i])
	}
	fmt.Println("___________________________________________________")
	for i := 0; i < 10; i++ {
		fmt.Println(dna.children.Memory[i])
	}
}
