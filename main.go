package main

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// ReceiptIDResponse represents the response for the receipt ID
type ReceiptIDResponse struct {
	ID string `json:"id"`
}

// PointsResponse represents the response for the points awarded
type PointsResponse struct {
	Points int64 `json:"points"`
}

var receipts = make(map[string]Receipt)

var (
	pricePattern    = regexp.MustCompile(`^\d+\.\d{2}$`)
	retailerPattern = regexp.MustCompile(`^[\w\s\-&]+$`)
)

func main() {
	r := gin.Default()

	r.POST("/receipts/process", processReceiptHandler)
	r.GET("/receipts/:id/points", getPointsHandler)

	r.Run(":8080")
}

func processReceiptHandler(c *gin.Context) {
	var receipt Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid."})
		return
	}
	if err := validateReceipt(receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid."})
		return
	}

	id := uuid.New().String()
	receipts[id] = receipt

	c.JSON(http.StatusOK, ReceiptIDResponse{ID: id})
}

func getPointsHandler(c *gin.Context) {
	id := c.Param("id")
	receipt, exists := receipts[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that ID."})
		return
	}

	points := calculatePoints(receipt)
	c.JSON(http.StatusOK, PointsResponse{Points: points})
}

func validateReceipt(receipt Receipt) error {
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

func calculatePoints(receipt Receipt) int64 {
	var points int64 = 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points++
		}
	}

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	total, err := parseTotal(receipt.Total)
	if err == nil && total == float64(int64(total)) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if total != 0 && int64(total*100)%25 == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	points += int64(len(receipt.Items) / 2 * 5)

	// Rule 5: If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedLength := len(item.ShortDescription)
		if trimmedLength%3 == 0 {
			price, err := parseTotal(item.Price)
			if err == nil {
				points += int64(price * 0.2)
			}
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd.
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}

func parseTotal(total string) (float64, error) {
	var value float64
	_, err := fmt.Sscanf(total, "%f", &value)
	if err != nil {
		return 0, errors.New("invalid total format")
	}
	return value, nil
}
