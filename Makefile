VERSION := $(shell git describe --tags --abbrev=0)

GREETING := helloworld
gen_opts := -t my_action user_booking

tag:
	@echo $(VERSION)
	@echo $(shell git describe --tags --abbrev=0) > VERSION

test:
	@go test -v ./...

gen:
	GREETING=${GREETING} go run cmd/gen/**.go generate ${gen_opts}

clear:
	go run cmd/gen/**.go clear ${gen_opts}


dry-run:
	GREETING=${GREETING} go run cmd/gen/**.go generate --dry-run ${gen_opts}

init:
	go run cmd/gen/**.go init
