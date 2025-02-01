package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PointsResponse struct {
	Points int64 `json:"points"`
}

func (h *Handler) GetPoints(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The ID is invalid: " + err.Error()})
	}
	receipt, exist := h.ReceiptService.GetReceipt(id)
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that ID."})
		return
	}

	c.JSON(http.StatusOK, PointsResponse{Points: h.PointService.CalculatePoints(receipt)})
}
