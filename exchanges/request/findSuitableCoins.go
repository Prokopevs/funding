package request

import (
	"funding/exchanges/types"
	"sort"
)

func findSuitableCoins(mainSlice *[]types.FundingItem, suitableCoinsSlice []types.SuitableCoin) (SliceRes []types.SuitableCoin) {
	const fundingPercent float64 = 0.05
	for _, exchange := range *mainSlice {
		for _, fundingElem := range exchange.FundingArr {
			fundingRate := fundingElem.FundingRate
			if fundingRate > fundingPercent|| fundingRate < -fundingPercent {

				detected := detectDublicate(fundingElem.Symbol, suitableCoinsSlice)
				if !detected {
					suitableCoinsSlice = append(suitableCoinsSlice, fundingElem)
				}
			}
		}
	}

	suitableCoinsSlice = sortSuitableSlice(&suitableCoinsSlice)
	return suitableCoinsSlice
}

func detectDublicate(symbol string, suitableCoinsSlice []types.SuitableCoin) (boo bool){
	var found = false;
	
	for i := 0; i < len(suitableCoinsSlice); i++ {
		if suitableCoinsSlice[i].Symbol == symbol {
			found = true
			break
		}
	}
	return found
}

func sortSuitableSlice(suitableCoinsSlice *[]types.SuitableCoin) (s []types.SuitableCoin) {
	positiveRates := make([]types.SuitableCoin, 0, 10)
	negativeRates := make([]types.SuitableCoin, 0, 10)


	for _, data := range *suitableCoinsSlice {
		if data.FundingRate >= 0 {
			positiveRates = append(positiveRates, data)
		} else {
			negativeRates = append(negativeRates, data)
		}
	}

	sort.Slice(positiveRates, func(i, j int) bool {
		return positiveRates[i].FundingRate > positiveRates[j].FundingRate
	})

	sort.Slice(negativeRates, func(i, j int) bool {
		return negativeRates[i].FundingRate < negativeRates[j].FundingRate
	})

	sortedArr := append(positiveRates, negativeRates...)

	return sortedArr
}