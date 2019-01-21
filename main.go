package main

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"net/http"
	"time"
	"os"
	"strconv"
	"fmt"
)

type echoRequestResource struct {
	Message string `json:"message"`
}

func (r echoRequestResource) validate() error {
	if r.Message == "" {
		return errors.New("\"message\" must not be empty")
	}

	return nil
}

type echoResponseResource struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type errorMessageResource struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	var data echoRequestResource
	var err error

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	err = decoder.Decode(&data)
	if err != nil {
		message := errorMessageResource{Code: 400, Message: "Malformed request"}
		w.WriteHeader(400)
		err = encoder.Encode(message)
		return
	}

	err = data.validate()
	if err != nil {
		message := errorMessageResource{Code: 400, Message: err.Error()}
		w.WriteHeader(400)
		err = encoder.Encode(message)
		return
	}

	response := echoResponseResource{Message: data.Message, Timestamp: time.Now().Format(time.RFC3339)}
	encoder.Encode(response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
}

func main() {
    port := os.Getenv("PORT")

    _, err := strconv.Atoi(port)
    if err != nil || port == "" {
        port = "5000"
    }

	r := chi.NewRouter()
	r.Post("/", echoHandler)
	r.Get("/", healthCheck)

	fmt.Printf("Server is listening on port %s\n", port)
	http.ListenAndServe(":" + port, r)
}
