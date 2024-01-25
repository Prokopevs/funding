package types

type BybitItem struct {
	Symbol          string `json:"symbol"`
	BidPrice        string `json:"bidPrice"`
	AskPrice        string `json:"askPrice"`
	FundingRate     string `json:"fundingRate"`
	NextFundingTime string `json:"nextFundingTime"`
}

type BybitResponse struct {
	Result struct {
		List []BybitItem `json:"list"`
	} `json:"result"`
}

type BitgetItem struct {
	Symbol      string `json:"symbol"`
	BidPr       string `json:"bidPr"`
	AskPr       string `json:"askPr"`
	FundingRate string `json:"fundingRate"`
}
type BitgetResponse struct {
	Data []BitgetItem `json:"data"`
}

type BinanceItem struct {
	Symbol          string `json:"symbol"`
	MarkPrice       string `json:"markPrice"`
	LastFundingRate string `json:"lastFundingRate"`
}

type MexcItem struct {
	Symbol      string `json:"symbol"`
	Bid1       float64 `json:"bid1"`
	Ask1       float64 `json:"ask1"`
	FundingRate float64 `json:"fundingRate"`
}
type MexcResponse struct {
	Data []MexcItem `json:"data"`
}

type GateIoItem struct {
	Contract        string `json:"contract"`
	Highest_bid     string `json:"highest_bid"`
	Lowest_ask       string `json:"lowest_ask"`
	Funding_rate string `json:"funding_rate"`
}
