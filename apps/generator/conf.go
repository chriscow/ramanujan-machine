package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type appConf struct {
	Constants []float64
	LHS       side
	RHS       side
}

func (c appConf) Save(filepath string) {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal("marshal:", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Getwd:", err)
	}

	err = ioutil.WriteFile(path.Join(cwd, filepath), b, os.FileMode(0755))
	if err != nil {
		log.Fatal("WriteFile:", err)
	}
}

func checkEnv() error {

	// TODO: verify required environment variables

	return nil
}

func loadArgs(argsFile string) error {

	// setDefaults()

	return nil
}
