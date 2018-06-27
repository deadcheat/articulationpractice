#! /usr/bin/make

.DEFAULT_GOAL := build

build:
	GOOS=linux go build twister.go
	-rm twister.zip > /dev/null
	zip twister.zip ./twister
