package main

import (
	"log"
	"math"
	"ramanujan/algorithm"
	"runtime"
	"strconv"
	"time"

	"golang.org/x/sync/semaphore"
)

const (
	defaultArgs = "default-args.json"
	argsUsage   = "-args args.json"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0)
	sem        = semaphore.NewWeighted(int64(maxWorkers))
)

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

	lhs := make(map[string][]algorithm.Solution)
	rhs := make(map[string][]algorithm.Solution)

	conf := bigConf()

	start := time.Now()
	process(conf.LHS, lhs)
	process(conf.RHS, rhs)
	elapsed := time.Since(start)

	log.Println("elapsed:", elapsed.Seconds())
	for k := range lhs {
		log.Println("lhs key:", k)
		if rlist, ok := rhs[k]; ok {
			// found in both sides
			log.Println("found", k, "on both sides")
			for i := range rlist {
				log.Println("index:", i, "type:", rlist[i].Type, "result:", rlist[i].Result)
			}
		}
	}
}
