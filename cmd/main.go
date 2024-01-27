package main

import (
	"funding/exchanges"
	"funding/internal/repository/funding"
	"funding/internal/services/funding"
	"funding/internal/transport/http/servers/funding/handler"
	"funding/internal/transport/http/servers/funding/router"
)


//  @title Funding API
//  @version 1.0
//	@description This is a sample funding server.

// @host localhost:8080
// @BasePath /
func main () {
	data := exchanges.NewFundingData()

	fundingRep := repository.NewfundingRepository(data)
	fundingSvc := services.NewService(fundingRep)
	fundingHandler := handler.NewHandler(fundingSvc)

	router.InitRouter(fundingHandler)

	router.Start("0.0.0.0:8080")
}