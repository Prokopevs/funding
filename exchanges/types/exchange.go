package types

type BybitItem struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	AskPrice string `json:"askPrice"`
	FundingRate string `json:"fundingRate"`
	NextFundingTime string `json:"nextFundingTime"`
}

type BybitResponse struct {
	Result  struct {List[]BybitItem `json:"list"`} `json:"result"`
}


type BitgetItem struct {
	Symbol   string `json:"symbol"`
	BidPr string `json:"bidPr"`
	AskPr string `json:"askPr"`
	FundingRate string `json:"fundingRate"`
}
type BitgetResponse struct {
	Data []BitgetItem `json:"data"` 
}