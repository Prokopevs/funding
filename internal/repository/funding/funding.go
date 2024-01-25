package repository

import (
	"funding/internal/model"
)
type FundingData interface {
	GetData() (*model.FinalData, error)
}