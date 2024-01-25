package main

import (
	// "encoding/json"
	// "fmt"
	"funding/exchanges"
	"funding/internal/repository/funding"
	"funding/internal/services/funding"
	"funding/internal/transport/http/servers/funding/handler"
	"funding/internal/transport/http/servers/funding/router"
)

func main () {
	data := exchanges.NewFundingData()

	fundingRep := repository.NewfundingRepository(data)
	fundingSvc := services.NewService(fundingRep)
	fundingHandler := handler.NewHandler(fundingSvc)

	router.InitRouter(fundingHandler)

	router.Start("0.0.0.0:8080")

	// result, _ :=  fundingSvc.GetData()
	// c, _ := json.Marshal(result)
    // fmt.Println(string(c))

}