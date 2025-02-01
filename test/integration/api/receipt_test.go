package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/dlccyes/receipt-processor/handler"
	"github.com/dlccyes/receipt-processor/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *Suite) TestProcessReceiptAndGetPoints() {
	testServer := httptest.NewServer(s.router)
	defer testServer.Close()

	processReceiptReq := &handler.ProcessReceiptReq{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []handler.ProcessReceiptReqItem{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			},
			{
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			},
			{
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			},
			{
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			},
			{
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            "12.00",
			},
		},
		Total: "35.35",
	}

	// save receipt
	req, err := http.NewRequest("POST", testServer.URL+"/receipts/process", test.ToRequestBody(processReceiptReq))
	require.NoError(s.T(), err)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(s.T(), err)
	defer resp.Body.Close()

	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)

	var responseBody handler.ProcessReceiptResp
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), handler.ProcessReceiptResp{
		ID: "1",
	}, responseBody)

	// get points
	req, err = http.NewRequest("GET", testServer.URL+"/receipts/"+responseBody.ID+"/points", nil)
	require.NoError(s.T(), err)

	resp, err = client.Do(req)
	assert.NoError(s.T(), err)
	defer resp.Body.Close()

	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)

	var pointsResponseBody handler.PointsResponse
	err = json.NewDecoder(resp.Body).Decode(&pointsResponseBody)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), handler.PointsResponse{
		Points: 28,
	}, pointsResponseBody)
}
