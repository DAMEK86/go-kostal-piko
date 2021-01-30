package app

import (
	"context"
	"fmt"
	"time"

	"github.com/damek86/go-kostal-piko/pkg/api"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	dataPointName = "pvwr"
	orgName       = ""
)

type Config struct {
	ScanInterval time.Duration
	InfluxDB     string
	ServerURL    string
	AuthToken    string
}

type KostalExporter struct {
	client       influxdb2.Client
	kostalClient api.Client
	cfg          Config
}

func NewKostalExporter(cfg Config, client api.Client) *KostalExporter {
	return &KostalExporter{
		client:       influxdb2.NewClient(cfg.ServerURL, cfg.AuthToken),
		cfg:          cfg,
		kostalClient: client,
	}
}

func (i *KostalExporter) Work() {
	fmt.Println("start executor")
	i.work()
	for true {
		select {
		case <-time.After(i.cfg.ScanInterval):
			i.work()
		}
	}
}

func (i *KostalExporter) work() {
	statsData := &api.DxsRespone{}
	err := i.kostalClient.GetStatsData(statsData)
	if err != nil {
		fmt.Println(err.Error())
	}

	generalData := &api.DxsRespone{}
	err = i.kostalClient.GetGeneralData(generalData)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Use blocking write client for writes to desired bucket
	writeAPI := i.client.WriteAPIBlocking(orgName, i.cfg.InfluxDB)
	dataPoints := make(map[string]interface{})
	for i := 0; i < len(statsData.Entries); i++ {
		dataPoints[statsData.Entries[i].FriendlyName] = statsData.Entries[i].Value
	}
	for i := 0; i < len(generalData.Entries); i++ {
		dataPoints[generalData.Entries[i].FriendlyName] = generalData.Entries[i].Value
	}
	p := influxdb2.NewPoint(dataPointName,
		nil,
		dataPoints, time.Now())
	err = writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Println(err)
		// do not handle error since writeAPI already print a message
	}
}
