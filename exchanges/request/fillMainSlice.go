package request

import (
	"funding/exchanges/types"
	"funding/exchanges/errW"
	"strings"
	"fmt"
	"encoding/json"
	"strconv"
	"errors"
)

func fillMainSlice(content *[]byte, urlReq string) (*types.FundingItem, error){

	if strings.Contains(urlReq, "bybit") {
		var response types.BybitResponse
		err := json.Unmarshal([]byte(*content), &response)
		if err != nil {
			errStr := fmt.Sprintf("Error Unmarshal data %s\n", err.Error())
			errW.ErrorHandler(errStr)
			return nil, err
		}

		localSlice := make([]types.SuitableCoin, 0, len(response.Result.List))
		
		for _, b:= range response.Result.List {
			fundingRate := convertToFloat(b.FundingRate) * 100
			formatFundingRate := convertToFloat(fmt.Sprintf("%.4f", fundingRate))

			newItem := types.SuitableCoin{
				Exchange: 		 "Bybit",
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
			return nil, err
		}
		
		localSlice := make([]types.SuitableCoin, 0, len(response.Data))
		
		for _, b:= range response.Data {
			fundingRate := convertToFloat(b.FundingRate) * 100
			formatFundingRate := convertToFloat(fmt.Sprintf("%.4f", fundingRate))
			

			newItem := types.SuitableCoin{
				Exchange: 		 "Bitget",
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