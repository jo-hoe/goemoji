include help.mk

# get root dir
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: test
test: ## runs all tests
	go test ./...

.PHONY: update-emojimap
update-emojimap: ## generates a new version of the emoji map
	go run $(ROOT_DIR)internal/main.go -output-path $(ROOT_DIR)emoji_map.json
