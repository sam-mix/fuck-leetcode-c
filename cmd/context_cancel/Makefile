.DEFAULT_GOAL := build

.PHONY: build run-ok run-cancel

build:
	go build -o context_cancel.exe .

run-ok: build
	./context_cancel.exe false

run-cancel: build
	./context_cancel.exe true
