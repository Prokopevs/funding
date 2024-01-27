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
	Symbol      string  `json:"symbol"`
	Bid1        float64 `json:"bid1"`
	Ask1        float64 `json:"ask1"`
	FundingRate float64 `json:"fundingRate"`
}
type MexcResponse struct {
	Data []MexcItem `json:"data"`
}

type GateIoItem struct {
	Contract     string `json:"contract"`
	Highest_bid  string `json:"highest_bid"`
	Lowest_ask   string `json:"lowest_ask"`
	Funding_rate string `json:"funding_rate"`
}

type BitmartItem struct {
	Symbol                string `json:"symbol"`
	Last_price            string `json:"last_price"`
	Expected_funding_rate string `json:"expected_funding_rate"`
}

type BitmartResponse struct {
	Data struct {
		Symbols []BitmartItem `json:"symbols"`
	} `json:"data"`
}

type BingxItem struct {
	Symbol          string `json:"symbol"`
	IndexPrice      string `json:"indexPrice"`
	LastFundingRate string `json:"lastFundingRate"`
}
type BingxResponse struct {
	Data []BingxItem `json:"data"`
}

type KucoinItem struct {
	Symbol         string  `json:"symbol"`
	IndexPrice     float64 `json:"indexPrice"`
	FundingFeeRate float64 `json:"fundingFeeRate"`
}
type KucoinResponse struct {
	Data []KucoinItem `json:"data"`
}

type CoinexItem struct {
	Symbol            string `json:"symbol"`
	Buy               string `json:"buy"`
	Sell              string `json:"sell"`
	Funding_rate_next string `json:"funding_rate_next"`
}

type CoinexResponse struct {
	Data struct {
		Ticker map[string]CoinexItem `json:"ticker"`
	} `json:"data"`
}

type HtxItem struct {
	Symbol       string `json:"symbol"`
	Fee_asset    string `json:"fee_asset"`
	Funding_rate interface{} `json:"funding_rate"`
}
type HtxResponse struct {
	Data []HtxItem `json:"data"`
}
