package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	serv FundingService
}

func NewHandler(s FundingService) *Handler {
	return &Handler{
		serv: s,
	}
}

func (r *Handler) GetData(c *gin.Context) {
	response, err := r.serv.GetData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}