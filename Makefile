test:
	unit_test=1 go test ./... -coverprofile=out/coverage.out -coverpkg=./... || true
	go tool cover -func=out/coverage.out
