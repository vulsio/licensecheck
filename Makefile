.PHONY: test

test:
	go test -cover ./license/...

generate-mock:
	mockgen -package mock -source license/shared/crawler.go -destination license/shared/mock/mock_crawler.go
