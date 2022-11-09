export PATH := $(GOPATH)/bin:$(PATH)
export GOPROXY := https://goproxy.cn
export PLUGINS_HOME := $(shell sh -c pwd)

define compile_plugin
	@echo "start compile plugin [$(1)] ..."
   	go build -buildmode=plugin -o $(PLUGINS_HOME)/plugins $(1)

endef

PLUGINS := $(shell ls -d ./plugins/*/)


build-plugin:
	@echo  $(PLUGINS)
	$(foreach plugin,$(PLUGINS),$(call compile_plugin,$(plugin)))

build-main:
	go build -o $(PLUGINS_HOME)/go-plugin ./main