package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"

	"golang.org/x/sync/semaphore"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0)
	sem        = semaphore.NewWeighted(int64(maxWorkers))
)

func (c appConf) Save() {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal("marshal:", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Getwd:", err)
	}

	err = ioutil.WriteFile(path.Join(cwd, ".conf/tiny.json"), b, os.FileMode(0755))
	if err != nil {
		log.Fatal("WriteFile:", err)
	}
}

func main() {

	if err := checkEnv(); err != nil {
		log.Fatal(err)
	}

	// the tiny config is just big enough to find e on both sides
	conf := tinyConf()

	// TODO: need to serialize args to solver that generated the value

	for v := range conf.LHS.Solve() {
		log.Println(v)
	}
}
