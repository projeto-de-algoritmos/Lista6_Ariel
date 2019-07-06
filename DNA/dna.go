package dna

import (
	"./layer1_alignmen-sequence"
	"./native_parents"
	"./native_tape"
	"./native_tapedist"
)

type DNA struct {
	allbuilded    bool
	verified      map[string]bool
	builded       map[string]bool
	guess_parents n_parents.parents
	children      n_tape.tape
	father        n_tape.tape
	distances     n_tapedist.tapedist
}
