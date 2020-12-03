package main

import (
	"flag"
	"log"
)

func main() {

	if err := checkEnv(); err != nil {
		log.Fatal(err)
	}

	var argsFile string
	flag.Parse()
	flag.StringVar(&argsFile, "args", defaultArgs, argsUsage)

	if err := loadArgs(argsFile); err != nil {
		log.Fatal("loadArgs:", err)
	}
}
