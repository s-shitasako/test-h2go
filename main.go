package main

import (
  "fmt"
  "net/http"
)

type H2Handler struct{}

var html []byte = ([]byte)("<h1>Go HTTP/2</h1>")

func (h *H2Handler)ServeHTTP(w http.ResponseWriter,r *http.Request) {
  fmt.Println(r.Proto, r.URL)

  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  w.Header().Set("Content-Length", fmt.Sprintf("%d", len(html)))
  w.Write(html)
}

func main() {
  handler := H2Handler{}
  http.ListenAndServeTLS(":4000", "cert.pem", "key.pem", &handler)
}
