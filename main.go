package main

import (
	"fmt"
	"net/http"
	"os"
)

type handler struct {
	name string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(fmt.Sprintf("Hello %v!", h.name)))
}

func main() {
	name := os.Getenv("NAME")

	if name == "" {
		name = "World"
	}

	err := http.ListenAndServe(":80", &handler{name})
	if err != nil {
		return
	}
}
