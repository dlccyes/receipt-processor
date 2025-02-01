package service_impl

import (
	"math"
	"strings"
	"unicode"

	"github.com/dlccyes/receipt-processor/model"
)

type pointServiceImpl struct{}

func (*pointServiceImpl) CalculatePoints(receipt *model.Receipt) int64 {
	points := int64(0)

	// #1 One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}

	// #2 50 points if the total is a round dollar amount with no cents.
	if math.Trunc(receipt.Total) == receipt.Total {
		points += 50
	}

	// #3 25 points if the total is a multiple of 0.25.
	if receipt.Total != 0 && int64(receipt.Total*100)%25 == 0 {
		points += 25
	}

	// #4 5 points for every two items on the receipt.
	points += int64(len(receipt.Items) / 2 * 5)

	// #5 If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedLen := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLen%3 != 0 {
			continue
		}
		points += int64(item.Price * 0.2)
	}

	// #6 6 points if the day in the purchase date is odd.
	if receipt.PurchaseDate.Day()%2 == 1 {
		points += 6
	}

	// #7 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if (receipt.PurchaseTime.Hour() >= 14 && receipt.PurchaseTime.Minute() >= 0) &&
		receipt.PurchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
