package main

import (
	"log"
	"time"

	"github.com/torwald-sergesson/app-a/pkg/client/v2"
)

const DefaultAddr = "localhost:8080"

func main() {
	cli := client.NewClient(DefaultAddr, time.Second*10)
	me, err := cli.Me()
	if err != nil {
		log.Fatalf("fail to get response: %s\n", err)
		return
	}
	log.Printf("Me: %#v\n", me)
}
