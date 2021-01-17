package main

import (
	"github.com/harvestcore/HarvestCCode/api"
	"github.com/harvestcore/HarvestCCode/log"
)

func main() {
	log.AddSimple(log.Info, "### HarvestCCode running ###")

	api := api.GetServer()
	api.Start()
}
