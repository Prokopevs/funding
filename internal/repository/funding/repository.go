package repository

import "funding/internal/model"

type FundingRepo struct {
	data FundingData
}

// NewfundingRepository create new repository
func NewfundingRepository(d FundingData) *FundingRepo {
	return &FundingRepo{
		data: d,
	}
}

func (r *FundingRepo) GetData() (*model.FinalData, error) {
	result, err := r.data.GetData()
	if err != nil {
		return nil, err
	}

	return result, nil
}

