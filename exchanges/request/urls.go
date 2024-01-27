package request

func getUrls() []string {
	urls := []string{
		"https://api.bybit.com/derivatives/v3/public/tickers",
		"https://api.bitget.com/api/v2/mix/market/tickers?productType=USDT-FUTURES",
		"https://fapi.binance.com/fapi/v1/premiumIndex",
		"https://contract.mexc.com/api/v1/contract/ticker",
		"https://api.gateio.ws/api/v4/futures/usdt/tickers",
		"https://api-cloud.bitmart.com/contract/public/details",
		"https://open-api.bingx.com/openApi/swap/v2/quote/premiumIndex",
		"https://api-futures.kucoin.com/api/v1/contracts/active",
		"https://api-cloud.bitmart.com/contract/public/details",
		"https://api.hbdm.com/linear-swap-api/v1/swap_batch_funding_rate",
	}

	return urls
}


