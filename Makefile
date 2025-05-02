GOROOT:=$(shell go env GOROOT)
DEVD:=$(shell command -v devd)

define DEV_ERROR
devd not found 😫.
Please install devd to serve the web files.
You can find it on github at https://github.com/cortesi/devd.git
endef

define HELP_MESSAGE

Usage: make [target]

Targets:
  all        - Build wasm and native binary (same as 'make web build')
  build      - Build native binary
  clean      - Clean up generated files
  help       - Show this help message
  run        - Runs the native binary
  serve      - Serve the web files using devd (must be installed)
  web        - Build only for the web (wasm)
endef

.PHONY: all web serve clean build run help

help:
	$(info $(HELP_MESSAGE))
	@echo

web/game.wasm: main.go
	@env GOOS=js GOARCH=wasm go build -o $@ .

web/wasm_exec.js: $(GOROOT)/lib/wasm/wasm_exec.js
	@cp $< $@

web/index.html:	resources/index.html
	@cp $< $@

web/game.html:	resources/game.html
	@cp $< $@

web: web/game.wasm web/wasm_exec.js web/index.html web/game.html
	@echo Publishing to the web...

bin/ebitentest: main.go
	@go build -o bin/ebitentest .
	
build: bin/ebitentest
	@echo Building ebitentest...

run: build
	@echo Running ebiten test...
	@./bin/ebitentest

serve: web
ifeq ($(DEVD),)
	$(error $(DEV_ERROR))
endif
	$(DEVD) -l -p 8080 -o -w ./web ./web

clean:
	@echo Removing 'web' folder
	@rm -rf web
	@echo Removing 'bin' folder
	@rm -rf bin

all: web build

.DEFAULT_GOAL := help
