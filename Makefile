# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: und android ios und-cross evm all test clean
.PHONY: und-linux und-linux-386 und-linux-amd64 und-linux-mips64 und-linux-mips64le
.PHONY: und-linux-arm und-linux-arm-5 und-linux-arm-6 und-linux-arm-7 und-linux-arm64
.PHONY: und-darwin und-darwin-386 und-darwin-amd64
.PHONY: und-windows und-windows-386 und-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

und:
	build/env.sh go run build/ci.go install ./cmd/und
	@echo "Done building."
	@echo "Run \"$(GOBIN)/und\" to launch und."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/und.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Und.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

lint: ## Run linters.
	build/env.sh go run build/ci.go lint

clean:
	./build/clean_go_build_cache.sh
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

und-cross: und-linux und-darwin und-windows und-android und-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/und-*

und-linux: und-linux-386 und-linux-amd64 und-linux-arm und-linux-mips64 und-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-*

und-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/und
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep 386

und-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/und
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep amd64

und-linux-arm: und-linux-arm-5 und-linux-arm-6 und-linux-arm-7 und-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep arm

und-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/und
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep arm-5

und-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/und
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep arm-6

und-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/und
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep arm-7

und-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/und
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep arm64

und-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/und
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep mips

und-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/und
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep mipsle

und-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/und
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep mips64

und-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/und
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/und-linux-* | grep mips64le

und-darwin: und-darwin-386 und-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/und-darwin-*

und-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/und
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/und-darwin-* | grep 386

und-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/und
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/und-darwin-* | grep amd64

und-windows: und-windows-386 und-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/und-windows-*

und-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/und
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/und-windows-* | grep 386

und-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/und
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/und-windows-* | grep amd64

devnet:
	docker-compose -f docker/docker-compose-mount.yml down --remove-orphans
	docker-compose -f docker/docker-compose-mount.yml up --build

devnet-down:
	docker-compose -f docker/docker-compose-mount.yml down --remove-orphans

devnet-pristine:
	docker-compose -f docker/docker-compose.yml down --remove-orphans
	docker-compose -f docker/docker-compose.yml up --build

devnet-pristine-down:
	docker-compose -f docker/docker-compose.yml down

devnet-syphoned-mount:
	docker-compose -f docker/docker-compose-syphoned-mount.yml down --remove-orphans
	docker-compose -f docker/docker-compose-syphoned-mount.yml up --build

devnet-syphoned-mount-down:
	docker-compose -f docker/docker-compose-syphoned-mount.yml down --remove-orphans
