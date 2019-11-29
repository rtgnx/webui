package main

import (
	"github.com/netauth/webui/internal/http"
)

func main() {
	s, err := http.New()
	if err != nil {
		panic(err)
	}

	s.Serve(":8080")
}
