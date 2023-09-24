.PHONY: build
build:
	@go build -o dist/podcastserver cmd/podcastserver/main.go

.PHONY: watch
watch:
	@air

.PHONY: clean
clean:
	@rm -fr dist
