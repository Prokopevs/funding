package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"funding/internal/model"
)

type Handler struct {
	serv FundingService
}

func NewHandler(s FundingService) *Handler {
	return &Handler{
		serv: s,
	}
}

// @Summary  	 Get funding list
// @Tags 		 funding
// @Description  get all funding
// @ID 			 get-funding
// @Accept 	 	 json
// @Produce 	 json
// @Success 	 200  {object} model.FinalData
// @Failure      500  {object}  model.ErrorResponse
// @Router       /funding [get]
func (r *Handler) GetData(c *gin.Context) {
	response, err := r.serv.GetData()
	if err != nil {
		c.JSON(http.StatusInternalServerError,  model.ErrorResponse{Err: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
