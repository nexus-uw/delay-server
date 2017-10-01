package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	maxMs, parseError := strconv.ParseInt(r.URL.Query().Get("max"), 0, 32)
	if parseError != nil {
		maxMs = 1
	}

	minMs, parseError := strconv.ParseInt(r.URL.Query().Get("min"), 0, 32)
	if parseError != nil {
		minMs = 0
	}

	failureChance, parseError := strconv.ParseInt(r.URL.Query().Get("failure"), 0, 32)
	if parseError != nil {
		failureChance = 0
	}

	if maxMs < 0 || maxMs > 60000 {
		http.Error(w, "invalid 'maxMs' query param. must be >= 0 and <= 60000", http.StatusBadRequest)
	}

	if minMs < 0 {
		http.Error(w, "invalid 'minMs' query param. must be >= 0", http.StatusBadRequest)
	}

	if maxMs < minMs {
		http.Error(w, "invalid 'maxMs' & 'minMs' query params. maxMs must be greater than or equal to minMs", http.StatusBadRequest)
	}

	time.Sleep(time.Duration(rand.Intn(int(maxMs-minMs))+int(minMs)) * time.Millisecond)

	if failureChance > 0 && rand.Intn(int(failureChance)) == 0 {
		http.Error(w, "Mock error", http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, "<h1>welcome to the go delay server</h1>"+
			"<h2> supported query params</h2>"+
			"<ul>"+
			"<li>maxMs : max delay in milliseconds. defaults to 1</li>"+
			"<li>minMs : min delay in milliseconds. defaults to 0</li>"+
			"<li>failure : failure chance. defaults to 0 (off)</li>"+
			"</ul>")
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
