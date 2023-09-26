.PHONY: build
build: clean
	@go build -o dist/radicaster cmd/radicaster/main.go

.PHONY: watch
watch:
	@air

.PHONY: clean
clean:
	@rm -fr dist
