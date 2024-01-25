package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"funding/exchanges/errW"
	"funding/exchanges/types"
	"strconv"
	"strings"
)

func fillMainSlice(content *[]byte, urlReq string) (*types.FundingItem, error) {

	if strings.Contains(urlReq, "bybit") {
		var response types.BybitResponse
		err := json.Unmarshal([]byte(*content), &response)
		if err != nil {
			errStr := fmt.Sprintf("Error Unmarshal data %s\n", err.Error())
			errW.ErrorHandler(errStr)
			fmt.Println(err)
			return nil, err
		}

		localSlice := make([]types.SuitableCoin, 0, len(response.Result.List))

		for _, b := range response.Result.List {
			fundingRate := convertToFloat(b.FundingRate) * 100
			formatFundingRate := convertToFloat(fmt.Sprintf("%.4f", fundingRate))

			newItem := types.SuitableCoin{
				Exchange:        "Bybit",
				Symbol:          b.Symbol,
				BidPrice:        b.BidPrice,
				AskPrice:        b.AskPrice,
				FundingRate:     formatFundingRate,
				NextFundingTime: b.NextFundingTime,
			}
			localSlice = append(localSlice, newItem)
		}

		obj := &types.FundingItem{Exchange: "Bybit", FundingArr: localSlice}

		return obj, nil
	}

	if strings.Contains(urlReq, "bitget") {
		var response types.BitgetResponse
		err := json.Unmarshal([]byte(*content), &response)
		if err != nil {
			errStr := fmt.Sprintf("Error Unmarshal data %s\n", err.Error())
			errW.ErrorHandler(errStr)
			fmt.Println(err)
			return nil, err
		}

		localSlice := make([]types.SuitableCoin, 0, len(response.Data))

		for _, b := range response.Data {
			fundingRate := convertToFloat(b.FundingRate) * 100
			formatFundingRate := convertToFloat(fmt.Sprintf("%.4f", fundingRate))

			newItem := types.SuitableCoin{
				Exchange:        "Bitget",
				Symbol:          b.Symbol,
				BidPrice:        b.BidPr,
				AskPrice:        b.AskPr,
				FundingRate:     formatFundingRate,
				NextFundingTime: "",
			}
			localSlice = append(localSlice, newItem)
		}

		obj := &types.FundingItem{Exchange: "Bitget", FundingArr: localSlice}

		return obj, nil
	}

	if strings.Contains(urlReq, "binance") {
		var response []types.BinanceItem
		err := json.Unmarshal([]byte(*content), &response)
		if err != nil {
			errStr := fmt.Sprintf("Error Unmarshal data %s\n", err.Error())
			errW.ErrorHandler(errStr)
			fmt.Println(err)
			return nil, err
		}

		localSlice := make([]types.SuitableCoin, 0, len(response))

		for _, b := range response {
			fundingRate := convertToFloat(b.LastFundingRate) * 100
			formatFundingRate := convertToFloat(fmt.Sprintf("%.4f", fundingRate))

			newItem := types.SuitableCoin{
				Exchange:        "Binance",
				Symbol:          b.Symbol,
				BidPrice:        b.MarkPrice,
				AskPrice:        b.MarkPrice,
				FundingRate:     formatFundingRate,
				NextFundingTime: "",
			}
			localSlice = append(localSlice, newItem)
		}

		obj := &types.FundingItem{Exchange: "Binance", FundingArr: localSlice}

		return obj, nil
	}

	if strings.Contains(urlReq, "mexc") {
		var response types.MexcResponse
		err := json.Unmarshal([]byte(*content), &response)
		if err != nil {
			errStr := fmt.Sprintf("Error Unmarshal data %s\n", err.Error())
			errW.ErrorHandler(errStr)
			fmt.Println(err)
			return nil, err
		}

		localSlice := make([]types.SuitableCoin, 0, len(response.Data))

		for _, b := range response.Data {
			fundingRate := b.FundingRate * 100
			formatFundingRate := convertToFloat(fmt.Sprintf("%.4f", fundingRate))

			formatSymbol := strings.ReplaceAll(b.Symbol, "_", "")

			newItem := types.SuitableCoin{
				Exchange:        "Mexc",
				Symbol:          formatSymbol,
				BidPrice:        strconv.FormatFloat(b.Bid1, 'f', -1, 64),
				AskPrice:        strconv.FormatFloat(b.Ask1, 'f', -1, 64),
				FundingRate:     formatFundingRate,
				NextFundingTime: "",
			}
			localSlice = append(localSlice, newItem)
		}

		obj := &types.FundingItem{Exchange: "Mexc", FundingArr: localSlice}

		return obj, nil
	}

	if strings.Contains(urlReq, "gateio") {
		var response []types.GateIoItem
		err := json.Unmarshal([]byte(*content), &response)
		if err != nil {
			errStr := fmt.Sprintf("Error Unmarshal data %s\n", err.Error())
			errW.ErrorHandler(errStr)
			fmt.Println(err)
			return nil, err
		}

		localSlice := make([]types.SuitableCoin, 0, len(response))

		for _, b := range response {
			fundingRate := convertToFloat(b.Funding_rate) * 100
			formatFundingRate := convertToFloat(fmt.Sprintf("%.4f", fundingRate))

			formatSymbol := strings.ReplaceAll(b.Contract, "_", "")

			newItem := types.SuitableCoin{
				Exchange:        "GateIo",
				Symbol:          formatSymbol,
				BidPrice:        b.Highest_bid,
				AskPrice:        b.Lowest_ask,
				FundingRate:     formatFundingRate,
				NextFundingTime: "",
			}
			localSlice = append(localSlice, newItem)
		}

		obj := &types.FundingItem{Exchange: "GateIo", FundingArr: localSlice}

		return obj, nil
	}

	return nil, errors.New("fail to find exchange")
}

func convertToFloat(s string) float64 {
	if s == "" {
		return 0
	}
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(s)

	}
	return n
}
