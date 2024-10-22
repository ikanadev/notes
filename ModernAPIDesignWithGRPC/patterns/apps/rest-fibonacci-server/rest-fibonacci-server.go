package restfibbonacciserver

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type App struct {
	asyncStores map[string]*AsyncStore
}

func NewApp() *App {
	return &App{
		asyncStores: make(map[string]*AsyncStore),
	}
}

func (app *App) Start() {
	router := mux.NewRouter()
	router.Use(generateRequestIDMiddleware)
	router.HandleFunc("/fibonacci/sync/{number:[0-9]+}", app.fibonacciSyncHandler)
	router.HandleFunc("/fibonacci/async/{number:[0-9]+}", app.fibonacciAsyncHandler)

	http.Handle("/", router)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func generateRequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("request-id")
		if strings.TrimSpace(reqID) == "" {
			r.Header.Add("request-id", uuid.New().String())
		}
		next.ServeHTTP(w, r)
	})
}
