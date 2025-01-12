package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		formattedTime := fmt.Sprintf("%s, %d %s %d", currentTime.Weekday(), currentTime.Day(), currentTime.Month(), currentTime.Year())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(formattedTime))
	})
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.Write([]byte("Hello there"))
		} else {
			w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
		}
	})
}

func main() {
	http.HandleFunc("/time", TimeHandler())
	http.HandleFunc("/hello", SayHelloHandler())
	http.ListenAndServe("localhost:8080", nil)
}
