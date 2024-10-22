package restfibonacciclient

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

var (
	typeOfCall string
	number     int
)

func init() {
	flag.StringVar(&typeOfCall, "typeOfCall", "sync", "do you want to call sync or async calls to server?")
	flag.IntVar(&number, "number", 10, "for what number do  you want  the fibonacci sequence")
}

type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) Start() {
	headerMap := map[string]string{"content-type": "application/json", "accept": "application/json"}
	baseURL := "http://localhost:8080"

	flag.Parse()

	if typeOfCall == "sync" {
		syncUrl := fmt.Sprintf("%s/fibonacci/sync/%d", baseURL, number)
		syncResult, err := sendRequest(syncUrl, headerMap)
		if err != nil {
			fmt.Printf("Error sending sync request: %s\n", err)
			return
		}
		fmt.Printf("sync fibo for %d = %s\n", number, syncResult)
	}
	if typeOfCall == "async" {
		for {
			asyncURL := fmt.Sprintf("%s/fibonacci/async/%d", baseURL, number)
			asyncResult, err := sendRequest(asyncURL, headerMap)
			if err != nil {
				fmt.Printf("Error sending async request: %s\n", err)
				return
			}

			endOfResponse := gjson.Get(asyncResult, "endOfResponse")

			if _, ok := headerMap["request-id"]; !ok {
				reqId := gjson.Get(asyncResult, "requestId")
				headerMap["request-id"] = reqId.String()
			}

			fmt.Printf("async fibo result for %d: %s\n", number, asyncResult)

			if endOfResponse.Exists() && endOfResponse.Bool() {
				break
			}

			time.Sleep(1 * time.Second)
		}
	}
}

func sendRequest(url string, headerMap map[string]string) (string, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	for k, v := range headerMap {
		req.Header.Add(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", fmt.Errorf("request failes with status code:  %d", response.StatusCode)
	}
	return string(body), nil
}
