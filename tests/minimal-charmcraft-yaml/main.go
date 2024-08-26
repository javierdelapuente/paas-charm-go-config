package main

import (
	"fmt"
	"go-app/config"
	"log"
	"net/http"
	"reflect"

	"github.com/caarlos0/env/v11"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	var cfg config.CharmConfig
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal("Error parsing configuration: %v", err)
	}

	metricsPort := 9001
	metricsPath := "/metrics"
	secretKey := "onerandomkey"
	httpProxy := "http://proxy.example.com:3128"
	expected := config.CharmConfig{
		ConfigOptions: config.ConfigOptions{
			Port:        9000,
			MetricsPort: &metricsPort,
			MetricsPath: &metricsPath,
			SecretKey:   &secretKey,
		},
		ProxyConfig: config.ProxyConfig{
			HTTPProxy:  &httpProxy,
			HTTPSProxy: nil,
			NoProxy:    []string{"127.0.0.1", "localhost", "::1"},
		},
	}
	log.Printf("Actual Config %+v\n", cfg)
	if !(reflect.DeepEqual(cfg, expected)) {
		log.Printf("Expected Config %+v\n", expected)
		log.Fatalf("Wrong configuration.")
	}

}
