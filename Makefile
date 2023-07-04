cover:
	go test -short -count=1 -race -coverprofile=coverage.out src/...
	go tool cover -html=coverage.out
	rm coverage.out

gen:
	/Users/ayto/go/bin/mockgen -source=src/order/usecase/interface.go \
            -destination=src/order/repository/mocks/mock_interface.go && \
    /Users/ayto/go/bin/mockgen -source=src/courier/usecase/interface.go \
                    -destination=src/courier/repository/mocks/mock_interface.go

.PHONY: gen cover