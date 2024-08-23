package main

import "flag"

var localMode bool
var prodEnv bool

func init() {
	flag.BoolVar(&localMode, "local", false, "Run locally instead of in a container")
	flag.BoolVar(&prodEnv, "prod", false, "Run in production environment")
	flag.Parse()
}

func main() {
	if localMode {
		runLocal()
	} else {
		runDockerized()
	}
}
