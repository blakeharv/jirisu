
package main

import (
  "log"
  "fmt"
  "net/http"
  "os"
)

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "<p>Blake's Bad Ass Go App of Supreme Awesomeness ('jirisu' is Japanese for 'gopher')</p>")
  fmt.Fprintln(w, "<p><a href=\"https://github.com/blakeharv/jirisu\">Repo Link</a></p>")
}

func main() {
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }
  http.HandleFunc("/", hello)
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
