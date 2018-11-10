package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GoToIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Oops, you're in the wrong place. To access the fibonacci API, navigate to `/api` endpoint.\n")
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n\nAvailable endpoints:\n`/api` -> Home\n`/api/fibonacci/:num` -> Returns Fibonacci sequence in JSON (ie: /api/fibonacci/6 returns {[0,1,1,2,3,5]})")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func fibonacciSeq(num int) []uint64 {
	if num == 0 {
		return []uint64{0}
	}

	if num == 1 {
		return []uint64{0}
	}

	sequence := []uint64{0, 1}
	for i := 2; i < num; i++ {
		sequence = append(sequence, sequence[i-1]+sequence[i-2])
	}
	return sequence
}

func FibonacciHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	num, err := strconv.Atoi(ps.ByName("num"))

	if err != nil {
		http.Error(w, "Please enter a valid number.", http.StatusInternalServerError)
		return
	}

	if num > 93 {
		http.Error(w, "Sequences greater than 93 are not allowed. (unit64 overflow)", http.StatusInternalServerError)
		return
	}

	fib, err := json.Marshal(fibonacciSeq(num))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(fib)
}

func main() {
	router := httprouter.New()
	router.GET("/", GoToIndex)
	router.GET("/api", Index)
	router.GET("/api/hello/:name", Hello)
	router.GET("/api/fibonacci/:num", FibonacciHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
