package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/damek86/go-kostal-piko/pkg/api"
	"net/http"
)

func main() {
	urlPtr := flag.String("url", "192.168.178.58", "remote address")
	usernamePtr := flag.String("username", api.DefaultUsername, "Kostal Piko username")
	passwordPtr := flag.String("password", api.DefaultPassword, "Kostal Piko password")

	flag.Parse()

	client := api.NewClient(http.DefaultClient, *urlPtr, *usernamePtr, *passwordPtr)
	work(client)
}

func work(client api.Client) {
	statsData := &api.DxsRespone{}
	err := client.GetStatsData(statsData)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", prettyPrint(statsData.Entries))

	generalData := &api.DxsRespone{}
	err = client.GetGeneralData(generalData)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", prettyPrint(generalData.Entries))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
