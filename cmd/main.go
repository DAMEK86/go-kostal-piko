package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/damek86/go-kostal-piko/pkg/api"
)

func main() {
	urlPtr := flag.String("url", "192.168.178.58", "remote address")
	usernamePtr := flag.String("username", api.DefaultUsername, "Kostal Piko username")
	passwordPtr := flag.String("password", api.DefaultPassword, "Kostal Piko password")
	scanIntervalPtr := flag.String("scan_interval", "30", "delay between request")

	flag.Parse()

	scanIntervalValue, err := strconv.Atoi(*scanIntervalPtr)
	if err != nil {
		panic(err)
	}
	scanInterval := time.Duration(scanIntervalValue) * time.Second
	client := api.NewClient(http.DefaultClient, *urlPtr, *usernamePtr, *passwordPtr)
	for true {
		work(client)
		time.Sleep(scanInterval)
	}
}

func work(client api.Client) {
	statsData := &api.DxsRespone{}
	err := client.GetStatsData(statsData)
	if err != nil {
		fmt.Println(err.Error())
	}

	generalData := &api.DxsRespone{}
	err = client.GetGeneralData(generalData)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", statsData.Entries)
}
