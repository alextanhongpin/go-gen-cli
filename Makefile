VERSION := $(shell git describe --tags --abbrev=0)

tag:
	@echo $(VERSION)
	#@echo $(shell git describe --tags --abbrev=0) > VERSION

test:
	@go test ./...

gen:
	PKG=user go run cmd/gen/**.go generate domain

clear:
	PKG=user go run cmd/gen/**.go clear domain
