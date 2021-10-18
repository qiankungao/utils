.PHONY:mod

mod:
	go mod download
	go mod tidy
	go mod vendor
