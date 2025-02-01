package handler

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/dlccyes/receipt-processor/model"
	"github.com/gin-gonic/gin"
)

type ReceiptIDResponse struct {
	ID string `json:"id"`
}

var (
	pricePattern    = regexp.MustCompile(`^\d+\.\d{2}$`)
	retailerPattern = regexp.MustCompile(`^[\w\s\-&]+$`)
)

func (h *Handler) ProcessReceipt(c *gin.Context) {
	var receipt model.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid: " + err.Error()})
		return
	}
	if err := validateReceipt(receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid: " + err.Error()})
		return
	}

	receiptID := h.ReceiptService.SaveReceipt(&receipt)

	c.JSON(http.StatusOK, ReceiptIDResponse{ID: receiptID})
}

func validateReceipt(receipt model.Receipt) error {
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || len(receipt.Items) == 0 || receipt.Total == "" {
		return errors.New("empty receipt")
	}
	if !retailerPattern.MatchString(receipt.Retailer) {
		return errors.New("invalid retailer format")
	}
	if !pricePattern.MatchString(receipt.Total) {
		return errors.New("invalid total format")
	}
	for _, item := range receipt.Items {
		if item.ShortDescription == "" || !pricePattern.MatchString(item.Price) {
			return errors.New("invalid price format")
		}
	}
	return nil
}
