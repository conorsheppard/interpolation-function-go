SHELL:=/bin/bash

default: build run

build:
	docker build -t interpolation-function-go .

run: build
	docker run --rm -t interpolation-function-go

.PHONY: build run