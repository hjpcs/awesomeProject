package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func GetFibonacciSeries(n int) []int {
	ret := make([]int, 2, n)
	ret[0] = 1
	ret[1] = 1
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome!"))
}

func createFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int
	for i := 0; i < 1000000; i++ {
		fbs = GetFibonacciSeries(50)
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs)))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/fb", createFBS)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

//cmdline
//127.0.0.1:9999/debug/pprof
//go tool pprof http://127.0.0.1:9999/debug/pprof/profile
//go-torch http://127.0.0.1:9999/debug/pprof/profile
