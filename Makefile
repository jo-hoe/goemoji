include help.mk

# get root dir
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))


.PHONY: update-emojimap
update-emojimap: ## generates a new version of the emoji map
	go run $(ROOT_DIR)internal/emojimap/main.go -output-path $(ROOT_DIR)goemoji/emoji_map.json