# Receipt Processor

## Details

This is a production grade repo with unnecessary dependency injection, layering, mocking, and full test coverage (unit test & integration test) for 2 simple backend APIs, written in Go. 

API requirements from <https://github.com/fetch-rewards/receipt-processor-challenge/>.

## How to run

Clone this repo

```
git clone https://github.com/dlccyes/receipt-processor.git
```

Go into the cloned repo

```
cd receipt-processor
```

Install dependencies

```
go mod tidy
```

Run the server

```
go run main.go
```

Save a receipt

```
curl --location 'http://localhost:8080/receipts/process' \
--header 'Content-Type: application/json' \
--data '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}'
```

Get the points of the receipt with the returned id

```
curl --location 'http://localhost:8080/receipts/1/points'
```
