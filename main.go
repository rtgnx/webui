package main

import (
	"os"

	"github.com/hashicorp/go-hclog"

	"github.com/netauth/webui/internal/http"
)

func main() {
	appLogger := hclog.New(&hclog.LoggerOptions{
		Name:  "webui",
		Level: hclog.LevelFromString("DEBUG"),
	})

	hclog.SetDefault(appLogger)

	s, err := http.New()
	if err != nil {
		appLogger.Error("Server initialization error", "error", err)
		os.Exit(1)
	}

	s.Serve(":8080")
}
