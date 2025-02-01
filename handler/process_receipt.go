package handler

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/dlccyes/receipt-processor/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type ProcessReceiptReq struct {
	Retailer     string                  `json:"retailer" binding:"required"`
	PurchaseDate string                  `json:"purchaseDate" binding:"required"`
	PurchaseTime string                  `json:"purchaseTime" binding:"required"`
	Items        []ProcessReceiptReqItem `json:"items" binding:"required"`
	Total        string                  `json:"total" binding:"required"`
}

type ProcessReceiptReqItem struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required"`
}

type ProcessReceiptResp struct {
	ID string `json:"id"`
}

var (
	retailerPattern = regexp.MustCompile(`^[\w\s\-&]+$`)
)

func (h *Handler) ProcessReceipt(c *gin.Context) {
	var req ProcessReceiptReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid: " + err.Error()})
		return
	}
	if err := validateReceipt(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid: " + err.Error()})
		return
	}
	receipt, err := toReceipt(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid: " + err.Error()})
	}
	receiptID := h.ReceiptService.SaveReceipt(receipt)

	c.JSON(http.StatusOK, ProcessReceiptResp{ID: strconv.FormatInt(receiptID, 10)})
}

func validateReceipt(receipt ProcessReceiptReq) error {
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || len(receipt.Items) == 0 || receipt.Total == "" {
		return errors.New("incomplete receipt")
	}
	if !retailerPattern.MatchString(receipt.Retailer) {
		return errors.New("invalid retailer format")
	}
	for _, item := range receipt.Items {
		if item.ShortDescription == "" {
			return errors.New("no short description")
		}
	}
	return nil
}

func toReceipt(req ProcessReceiptReq) (*model.Receipt, error) {
	purchaseDate, err := time.Parse("2006-01-02", req.PurchaseDate)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid purchase date")
	}
	purchaseTime, err := time.Parse("15:04", req.PurchaseTime)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid purchase time")
	}
	total, err := strconv.ParseFloat(req.Total, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid total format")
	}
	items, err := toReceiptItem(req.Items)
	if err != nil {
		return nil, err
	}
	return &model.Receipt{
		Retailer:     req.Retailer,
		PurchaseDate: purchaseDate,
		PurchaseTime: purchaseTime,
		Items:        items,
		Total:        total,
	}, nil
}

func toReceiptItem(items []ProcessReceiptReqItem) ([]model.Item, error) {
	modelItems := make([]model.Item, 0, len(items))
	for _, item := range items {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid price format")
		}
		modelItems = append(modelItems, model.Item{
			ShortDescription: item.ShortDescription,
			Price:            price,
		})
	}
	return modelItems, nil
}
