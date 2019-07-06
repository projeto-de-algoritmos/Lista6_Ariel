//package main

package n_tape

type string_reader interface {
	readlines(path string) error
	//verifydata(path string) ([]string, error)
}
