.PHONY: build
build:
	@go build -o dist/podcastserver cmd/podcastserver/main.go

.PHONY: clean
clean:
	@rm -fr dist