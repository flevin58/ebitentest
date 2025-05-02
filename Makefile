GOROOT:=$(shell go env GOROOT)
DEVD:=$(shell command -v devd)

define DEV_ERROR
devd not found 😫.
Please install devd to serve the web files.
You can find it on github at https://github.com/cortesi/devd.git
endef

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

serve: web
ifeq ($(DEVD),)
	$(error $(DEV_ERROR))
endif
	$(DEVD) -l -p 8080 -o -w ./web ./web

clean:
	@echo Removing 'web' folder
	@rm -rf web

all: web

.PHONY: all web serve clean
