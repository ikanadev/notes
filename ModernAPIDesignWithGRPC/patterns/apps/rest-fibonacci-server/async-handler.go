package restfibbonacciserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

type AsyncResponse struct {
	RequestId        string `json:"requestId"`
	FibonacciNumbers []int  `json:"fibonacciNumbers"`
	EndOfResponse    bool   `json:"endOfResponse"`
}

type AsyncStore struct {
	mu             sync.Mutex
	current        int
	requestedRange int
	numbers        []int
}

func NewAsyncStore(requestedRange int) *AsyncStore {
	return &AsyncStore{requestedRange: requestedRange}
}

func (store *AsyncStore) Write(number, current int) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.current = current
	store.numbers = append(store.numbers, number)
}

func (store *AsyncStore) Read() ([]int, int, int) {
	store.mu.Lock()
	defer store.mu.Unlock()
	readNumbers := make([]int, len(store.numbers))
	copy(readNumbers, store.numbers)
	store.numbers = []int{}
	return readNumbers, store.current, store.requestedRange
}

func (app *App) fibonacciAsyncHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["number"]

	numFibonacci, err := strconv.Atoi(number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	headers := r.Header
	reqId := headers.Get("request-id")
	if strings.TrimSpace(reqId) == "" {
		http.Error(w, "request-id not found in headers", http.StatusBadRequest)
		return
	}

	if _, ok := app.asyncStores[reqId]; !ok {
		fmt.Printf("creating new async store for %s\n", reqId)
		app.asyncStores[reqId] = NewAsyncStore(numFibonacci)
		go app.fibAsync(numFibonacci, reqId)
	}

	numbersNow, current, requested := app.asyncStores[reqId].Read()
	fmt.Printf("read fibs reqId %s till current %d and number are : %v\n", reqId, current, numbersNow)
	end := false
	if current == requested {
		end = true
		delete(app.asyncStores, reqId)
	}
	response := AsyncResponse{
		RequestId:        reqId,
		FibonacciNumbers: numbersNow,
		EndOfResponse:    end,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func (app *App) fibAsync(n int, reqId string) {
	for i := 0; i < n; i++ {
		fmt.Printf("for %s computing and writing fib  of %d\n", reqId, i)
		app.asyncStores[reqId].Write(fib(i), i+1)
	}
}
