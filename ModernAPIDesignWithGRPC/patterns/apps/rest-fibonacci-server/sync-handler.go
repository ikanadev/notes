package restfibbonacciserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type SyncResponse struct {
	TimeTaken        string `json:"timeTaken"`
	FibonacciNumbers []int  `json:"fibonacciNumbers"`
}

func (app *App) fibonacciSyncHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["number"]

	numFibonacci, err := strconv.Atoi(number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fibonacciNumbers := make([]int, numFibonacci)
	now := time.Now()

	for i := 0; i < numFibonacci; i++ {
		fibonacciNumbers[i] = fib(i)
	}

	timeTaken := time.Since(now).Milliseconds()
	response := SyncResponse{
		TimeTaken:        fmt.Sprintf("%dms", timeTaken),
		FibonacciNumbers: fibonacciNumbers,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
