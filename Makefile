cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

gen:
	mockgen -source=internal/order/order.go \
            -destination=internal/repository/order/mocks/mock_interface.go && \
    mockgen -source=internal/courier/courier.go \
                    -destination=internal/repository/courier/mocks/mock_interface.go

.PHONY: gen cover