package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"ramanujan/algorithm"
	"runtime"
	"strconv"

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

func process(s side, m map[string][]algorithm.Solution) {
	for v := range s.Solve() {
		// the fractional part is the key
		_, frac := math.Modf(v.Result)
		key := strconv.FormatFloat(frac, 'f', 8, 64)[2:]
		// key := strconv.FormatFloat(v.Result, 'f', 8, 64)
		if _, ok := m[key]; !ok {
			res := make([]algorithm.Solution, 0)
			m[key] = res
		}

		// for k, v := range m[key] {
		// TODO: look for duplicates using Solution::Hash
		// }

		m[key] = append(m[key], v)
	}
}

func main() {

	if err := checkEnv(); err != nil {
		log.Fatal(err)
	}
	// the tiny config is just big enough to find e on both sides

	// TODO: need to serialize args to solver that generated the value

}
