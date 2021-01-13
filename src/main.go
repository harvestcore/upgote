package main

import (
	"github.com/harvestcore/HarvestCCode/src/api"
	"github.com/harvestcore/HarvestCCode/src/log"
)

func main() {
	log.AddSimple(log.Info, "### HarvestCCode running ###")

	api := api.GetServer()
	api.Start()
}
