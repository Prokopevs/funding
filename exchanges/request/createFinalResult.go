package request

import (
	"funding/exchanges/types"
	"funding/internal/model"
)

func createFinalResult(mainSlice *[]types.FundingItem, suitableCoinsSlice *[]types.SuitableCoin) (res model.FinalData) {
	var result model.FinalData 
	positiveArray := make([]types.FundingCoin, 0, 20)
	negativeArray := make([]types.FundingCoin, 0, 20)

	for _, item := range *suitableCoinsSlice {
		arr := make([]types.SuitableCoin, 0, 10)
		arr = append(arr, item)
		symbol := item.Symbol
		exchange := item.Exchange
		fundingRate := item.FundingRate

		for _, exchan := range *mainSlice {
			if exchange == exchan.Exchange {
				continue
			}
			for _, fundingElem := range exchan.FundingArr {
				if symbol == fundingElem.Symbol {
					arr = append(arr, fundingElem)
				}
			}
		}

		obj := types.FundingCoin{Symbol: symbol, Elems: arr}
		if fundingRate > 0 {
			positiveArray = append(positiveArray, obj)
		} else {
			negativeArray = append(negativeArray, obj)
		}
	}

	result.Positive = positiveArray
	result.Negative = negativeArray
	
	return result
}