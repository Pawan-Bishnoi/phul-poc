package main
import (
  "fmt"
  "net/http"
  "phul-poc/runner"
  "phul-poc/definition"
)
func main() {

  http.HandleFunc("/health", health)
  fmt.Println("Core is up")
  ch := make(chan definition.Message, 1)
  runner.Init(ch)
  http.ListenAndServe(":8000", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "alive")
}
