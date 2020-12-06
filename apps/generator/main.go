package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
)

type appConf struct {
	Constants []float64
	LHS       Side
	RHS       Side
}

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
	lhs, rhs := tiny()
	conf := appConf{
		Constants: []float64{math.E},
		LHS:       lhs,
		RHS:       rhs,
	}

	for v := range conf.LHS.Solve() {
		log.Println(v)
	}
}
