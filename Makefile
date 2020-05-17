# (C) Copyright 2019 Mario Trangoni
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

GO                      ?= GO111MODULE=on go
GOPATH                  := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
GOOPTS                  := $(GOOPTS) -mod=vendor
GOLINTER                ?= $(GOPATH)/bin/golangci-lint
pkgs                    = $(shell $(GO) list ./... | grep -v /vendor/)
TARGET                  ?= geoip2_zabbix

PREFIX                  ?= $(shell pwd)
BIN_DIR                 ?= $(shell pwd)

.PHONY: all
all: clean common-deps format vet golangci build test

.PHONY: test
test:
	@echo ">> running tests"
	@$(GO) test -v $(pkgs)

.PHONY: format
format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

.PHONY: vet
vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

.PHONY: golangci
golangci: $(GOLINTER)
	@echo ">> linting code"
	@$(GOLINTER) run --config ./.golanci.yml

.PHONY: build
build: common-deps
	@echo ">> building binaries"
	@$(GO) build -tags netgo -ldflags '-s -extldflags "-static"' -o $(TARGET) $(pkgs)

.PHONY: crossbuild
crossbuild: common-deps
	@echo ">> crossbuilding binaries for windows"
	@GOOS=windows GOARCH=amd64 $(GO) build -tags netgo -ldflags '-s -extldflags "-static"' -o $(TARGET).exe $(pkgs)

.PHONY: clean
clean:
	@echo ">> Cleaning up"
	@find . -type f -name '*~' -exec rm -fv {} \;
	@$(RM) $(TARGET)

.PHONY: common-deps
common-deps:
	@echo ">> ensure vendoring"
	@$(GO) mod download

.PHONY: download-geolite2-city
download-geolite2-city: geoipupdate
	@echo ">> download GeoLite2-City DB"
	@geoipupdate --config-file ./GeoIP.conf -d .

.PHONY: golangci-lint lint
$(GOPATH)/bin/golangci-lint lint:
	@GOOS=$(shell uname -s | tr A-Z a-z) \
		GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m))) \
		$(GO) get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0

.PHONY: geoipupdate
$(GOPATH)/bin/geoipupdate geoipupdate:
	@GOOS=$(shell uname -s | tr A-Z a-z) \
		GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m))) \
		$(GO) get github.com/maxmind/geoipupdate/v4/cmd/geoipupdate
