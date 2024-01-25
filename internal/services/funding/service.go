package services

import "funding/internal/model"
type serv struct {
	repo    FundingRepo
}

func NewService(r FundingRepo) *serv {
	return &serv{
		repo: r,
	}
}

func (r *serv) GetData() (*model.FinalData, error) {
	result, err := r.repo.GetData()
	if err != nil {
		return nil, err
	}

	return result, nil
}