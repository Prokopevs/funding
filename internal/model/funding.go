package model

import (
	"funding/exchanges/types"
	
)

type FinalData struct {
	Positive []types.FundingCoin `json:"positive"`
	Negative []types.FundingCoin `json:"negative"`
}

type ErrorResponse struct {
	Err string `json:"err"`
}