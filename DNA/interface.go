package main

type testbuilder interface {
	setupfathers(path string) error
	setupchildren(path string) error
	//verifydata(path string) ([]string, error)
}
