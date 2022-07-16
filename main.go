package main

import (
	"github.com/harvestcore/upgote/api"
	"github.com/harvestcore/upgote/log"
)

func main() {
	log.AddSimple(log.Info, "UPGOTE is now running.")

	api := api.GetServer()
	api.Start()
}
