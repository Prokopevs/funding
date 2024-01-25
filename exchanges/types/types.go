package types

type FundingItem struct {
	Exchange    string 
	FundingArr  []SuitableCoin
}

type SuitableCoin struct {
	Exchange string `json:"exchange"`
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	AskPrice string `json:"askPrice"`
	FundingRate float64 `json:"fundingRate"`
	NextFundingTime string `json:"nextFundingTime"`
}

type FundingCoin struct {
	Symbol    string `json:"symbol"`
	Elems  []SuitableCoin `json:"elems"`
}

