cover:
	go test -short -count=1 -race -coverprofile=coverage.out src/...
	go tool cover -html=coverage.out
	rm coverage.out

gen:
	mockgen -source=src/order/order.go \
            -destination=src/order/mocks/mock_interface.go && \
    mockgen -source=src/courier/courier.go \
                    -destination=src/courier/mocks/mock_interface.go

.PHONY: gen cover