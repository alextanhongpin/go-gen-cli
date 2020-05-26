tag:
	@echo $(shell git describe --tags --abbrev=0) > VERSION
