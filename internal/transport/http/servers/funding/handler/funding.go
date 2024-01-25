package handler


import (
	"funding/internal/model"
)

type FundingService interface {
	GetData() (*model.FinalData, error)
}