package request

func getUrls() ([]string) {
	urls := []string{
		"https://api.bybit.com/derivatives/v3/public/tickers",
		"https://api.bitget.com/api/v2/mix/market/tickers?productType=USDT-FUTURES",
	}

	return urls
}