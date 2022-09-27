package main

import (
	"strconv"
	"sync"

	"net/http"
	_ "net/http/pprof"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	var s []string
	for i := 0; i < 1000; i++ {
		s = append(s, strconv.Itoa(i))
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(&wg)
	go func() {
		http.ListenAndServe(":8000", nil)
	}()
	wg.Wait()
}
