package handler

import (
	"net/http"

	"github.com/dlccyes/receipt-processor/service"
	"github.com/gin-gonic/gin"
)

type PointsResponse struct {
	Points int64 `json:"points"`
}

func (h *Handler) GetPoints(c *gin.Context) {
	receipt, exist := service.GetReceipt(c.Param("id"))
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that ID."})
		return
	}

	c.JSON(http.StatusOK, PointsResponse{Points: service.CalculatePoints(receipt)})
}
