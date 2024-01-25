package exchanges

import (
	"fmt"
	"funding/exchanges/request"
	"funding/exchanges/types"
	"time"
	"funding/internal/model"
	"sync"
)

type FundingData struct {
	Data model.FinalData
	Deadline time.Time
	Mu sync.RWMutex
}

func NewFundingData() *FundingData {
	return &FundingData {
		Data: model.FinalData{Positive: make([]types.FundingCoin, 0), Negative: make([]types.FundingCoin, 0)},
		Deadline: time.Now().Add(30 * time.Second),
		Mu: sync.RWMutex{},
	}
}

func (s *FundingData) GetData() (*model.FinalData, error) {
	now := time.Now()
	
	if s.Deadline.After(now) && (len(s.Data.Positive) != 0 || len(s.Data.Negative) != 0) {
		s.Mu.RLock() 
		defer s.Mu.RUnlock()
		fmt.Println("cashed")

		dataCopy := s.Data // Копирование структуры
		return &dataCopy, nil
    } else {
		s.Mu.Lock()
		defer s.Mu.Unlock()

		result, err := request.GetFunding()
		if err != nil {
			return nil, err
		}

		s.Data = *result
		s.Deadline = time.Now().Add(30 * time.Second)

		return result, nil
	}
}