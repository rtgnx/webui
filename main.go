package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/netauth/webui/internal/http"
)

var (
	cfg = pflag.String("config", "", "Config file")
)

func main() {
	appLogger := hclog.New(&hclog.LoggerOptions{
		Name:  "webui",
		Level: hclog.LevelFromString("DEBUG"),
	})
	hclog.SetDefault(appLogger)

	viper.BindPFlags(pflag.CommandLine)
	if *cfg != "" {
		viper.SetConfigFile(*cfg)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.netauth")
		viper.AddConfigPath("/etc/netauth/")
	}
	if err := viper.ReadInConfig(); err != nil {
		appLogger.Error("Error reading config", "error", err)
		os.Exit(1)
	}
	viper.Set("client.ServiceName", "webui")

	s, err := http.New()
	if err != nil {
		appLogger.Error("Server initialization error", "error", err)
		os.Exit(1)
	}

	if err := s.Serve(":8080"); err != nil {
		appLogger.Error("Error initializing webserver", "error", err)
		os.Exit(1)
	}
}
