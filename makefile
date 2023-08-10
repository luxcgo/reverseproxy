SHELL := /bin/bash

build:
	docker build \
		-f dockerfile \
		-t reverseproxy:latest \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.