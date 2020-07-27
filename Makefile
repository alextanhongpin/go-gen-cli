VERSION := $(shell git describe --tags --abbrev=0)

tag:
	@echo $(VERSION)
	#@echo $(shell git describe --tags --abbrev=0) > VERSION

test:
	@go test -v ./...

gen:
	PKG=user go run cmd/gen/**.go generate -t domain user_booking

clear:
	PKG=user go run cmd/gen/**.go clear -t domain user_booking


dry-run:
	PKG=user go run cmd/gen/**.go generate --dry-run -t domain user_booking
