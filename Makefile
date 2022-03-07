.PHONY: test

test:
	go test -cover ./...

generate-mock:
	mockgen -package mock -source license/shared/crawler.go -destination license/shared/mock/mock_crawler.go
