package services

import (
	"funding/internal/model"
)

type FundingRepo interface {
	GetData() (*model.FinalData, error)
}