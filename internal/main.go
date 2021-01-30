package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/damek86/go-kostal-piko/internal/app"
	"github.com/damek86/go-kostal-piko/pkg/api"
)

func main() {
	client := api.NewClient(http.DefaultClient,
		getEnv("KOSTAL_URL", "192.168.178.58"),
		getEnv("KOSTAL_USERNAME", api.DefaultUsername),
		getEnv("KOSTAL_PASSWORD", api.DefaultPassword),
	)
	go app.StartHealthEndpoint()
	app.NewKostalExporter(getAppConfig(), client).Work()
}

func getAppConfig() app.Config {
	scanInterval, err := strconv.Atoi(getEnv("SCAN_INTERVAL", "30"))
	if err != nil {
		panic(err)
	}
	return app.Config{
		ScanInterval: time.Duration(scanInterval) * time.Second,
		InfluxDB:     getEnv("INFLUXDB_DB", "solar"),
		ServerURL:    fmt.Sprintf("%s:%s", getEnv("INFLUXDB_HOST", "http://localhost"), getEnv("INFLUXDB_PORT", "8086")),
		AuthToken:    fmt.Sprintf("%s:%s", getEnv("INFLUXDB_USER", "root"), getEnv("INFLUXDB_PASSWORD", "root")),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
