package main

import (
	"fmt"
	"net/http"
	"phul-poc/definition"
	"phul-poc/runner"
)

func main() {

	http.HandleFunc("/health", health)
	fmt.Println("Core is up")
	ch := make(chan definition.Message, 1)
	runner.Init(ch)
	for true {
		fmt.Printf("received over channel %v", <-ch)
	}
	// http.ListenAndServe(":8000", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "alive")
}
