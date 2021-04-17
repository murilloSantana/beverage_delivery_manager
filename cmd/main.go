package main

import (
	"beverage_delivery_manager/cmd/server"
	"beverage_delivery_manager/cmd/settings"
	"log"
)

func main() {
	sts := settings.New()

	log.Fatal(server.New(sts))
}
