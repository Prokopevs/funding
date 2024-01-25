package request

import (
	"fmt"
	"funding/exchanges/errW"
	"funding/exchanges/types"
	"funding/internal/model"
	"io"
	"net/http"
	"sync"
)

func GetFunding() (*model.FinalData, error) {
	urls := getUrls()

	mainSlice := make([]types.FundingItem, 0, len(urls))
	suitableCoinsSlice := make([]types.SuitableCoin, 0, 30)
	var result model.FinalData

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, u := range urls {

		wg.Add(1)
		go func(url string) {

			defer wg.Done()

			content, urlReq := doReq(&url)
			if content != nil {
				mainSliceElem, err := fillMainSlice(content, urlReq)
				if err == nil {
					mu.Lock()
					mainSlice = append(mainSlice, *mainSliceElem)
					mu.Unlock()
				}
			}
		}(u)
	}
	wg.Wait()

	suitableCoinsSlice = findSuitableCoins(&mainSlice, suitableCoinsSlice)
	result = createFinalResult(&mainSlice, &suitableCoinsSlice)

	return &result, nil
}

func doReq(url *string) (content *[]byte, urlReq string) {
	resp, err := http.Get(*url)
	if err != nil {
		errStr := fmt.Sprintf("Error requesting data from %s: %s\n", *url, err.Error())
		errW.ErrorHandler(errStr)
		fmt.Println(err)

		return nil, *url
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errStr := fmt.Sprintf("Error reading body from %s: %s\n", *url, err.Error())
		errW.ErrorHandler(errStr)
		fmt.Println(err)
		
		return nil, *url
	}

	return &body, *url
}
