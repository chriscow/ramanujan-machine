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
	LHS       SideConf
	RHS       SideConf
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

	lhs, rhs := tiny()
	conf := appConf{
		Constants: []float64{math.E, math.Phi},
		LHS:       lhs,
		RHS:       rhs,
	}

	for v := range conf.LHS.Solve() {
		log.Println(v)
	}
}
