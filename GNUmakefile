.DEFAULT_GOAL := build

GOPATH ?= $(HOME)/go
COMMIT_SHA ?= $(shell git rev-parse HEAD)
VERSION = $(shell cat VERSION)
GOBIN ?= $(GOPATH)/bin
GO_LDFLAGS := $(shell tools/bin/ldflags)
GOFLAGS = -ldflags "$(GO_LDFLAGS)"

.PHONY: install
install: operator-ui-autoinstall install-plugin-autoinstall ## Install plugin and all its dependencies.

.PHONY: install-git-hooks
install-git-hooks: ## Install git hooks.
	git config core.hooksPath .githooks

.PHONY: install-plugin-autoinstall
install-plugin-autoinstall: | pnpmdep gomod install-plugin ## Autoinstall plugin.
.PHONY: operator-ui-autoinstall
operator-ui-autoinstall: | operator-ui ## Autoinstall frontend UI.

.PHONY: pnpmdep
pnpmdep: ## Install solidity contract dependencies through pnpm
	(cd contracts && pnpm i)
	
.PHONY: gomod
gomod: ## Ensure chainlink's go dependencies are installed.
	@if [ -z "`which gencodec`" ]; then \
		go install github.com/smartcontractkit/gencodec@latest; \
	fi || true
	go mod download

.PHONY: gomodtidy
gomodtidy: ## Run go mod tidy on all modules.
	go mod tidy
	cd ./integration-tests && go mod tidy

.PHONY: install-plugin
install-plugin: plugin ## Install the plugin binary.
	mkdir -p $(GOBIN)
	rm -f $(GOBIN)/plugin
	cp $< $(GOBIN)/plugin

plugin: operator-ui ## Build the plugin binary.
	go build -o $@ ./core/
	#go build $(GOFLAGS) -o $@ ./core/

.PHONY: docker ## Build the plugin docker image
docker:
	docker buildx build \
	--build-arg COMMIT_SHA=$(COMMIT_SHA) \
	-f core/plugin.Dockerfile .

.PHONY: plugin-build
plugin-build: operator-ui ## Build & install the plugin binary.
	go build -o plugin ./core/
	rm -f $(GOBIN)/plugin
	cp plugin $(GOBIN)/plugin

.PHONY: operator-ui
operator-ui: ## Fetch the frontend
	./operator_ui/install.sh

.PHONY: abigen
abigen: ## Build & install abigen.
	./tools/bin/build_abigen

.PHONY: go-solidity-wrappers
go-solidity-wrappers: pnpmdep abigen ## Recompiles solidity contracts and their go wrappers.
	./contracts/scripts/native_solc_compile_all
	go generate ./core/gethwrappers

.PHONY: go-solidity-wrappers-ocr2vrf
go-solidity-wrappers-ocr2vrf: pnpmdep abigen ## Recompiles solidity contracts and their go wrappers.
	./contracts/scripts/native_solc_compile_all_ocr2vrf
	# replace the go:generate_disabled directive with the regular go:generate directive
	sed -i '' 's/go:generate_disabled/go:generate/g' core/gethwrappers/ocr2vrf/go_generate.go
	go generate ./core/gethwrappers/ocr2vrf
	go generate ./core/internal/mocks
	# put the go:generate_disabled directive back
	sed -i '' 's/go:generate/go:generate_disabled/g' core/gethwrappers/ocr2vrf/go_generate.go

.PHONY: generate
generate: abigen ## Execute all go:generate commands.
	go generate -x ./...

.PHONY: testdb
testdb: ## Prepares the test database.
	go run ./core/main.go local db preparetest

.PHONY: testdb
testdb-user-only: ## Prepares the test database with user only.
	go run ./core/main.go local db preparetest --user-only

# Format for CI
.PHONY: presubmit
presubmit: ## Format go files and imports.
	goimports -w ./core
	gofmt -w ./core
	go mod tidy

.PHONY: mockery
mockery: $(mockery) ## Install mockery.
	go install github.com/vektra/mockery/v2@v2.20.0

.PHONY: telemetry-protobuf
telemetry-protobuf: $(telemetry-protobuf) ## Generate telemetry protocol buffers.
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-wsrpc_out=. \
	--go-wsrpc_opt=paths=source_relative \
	./core/services/synchronization/telem/*.proto

.PHONY: test_need_operator_assets
test_need_operator_assets: ## Add blank file in web assets if operator ui has not been built
	[ -f "./core/web/assets/index.html" ] || mkdir ./core/web/assets && touch ./core/web/assets/index.html

.PHONY: config-docs
config-docs: ## Generate core node configuration documentation
	go run ./core/config/v2/docs/cmd/generate/main.go -o ./docs/

.PHONY: golangci-lint
golangci-lint: ## Run golangci-lint for all issues.
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run --max-issues-per-linter 0 --max-same-issues 0 > golangci-lint-output.txt

.PHONY: snapshot
snapshot:
	cd ./contracts && forge snapshot --match-test _gas

GORELEASER_CONFIG ?= .goreleaser.yaml

.PHONY: goreleaser-dev-build
goreleaser-dev-build: ## Run goreleaser snapshot build
	./tools/bin/goreleaser_wrapper build --snapshot --rm-dist --config ${GORELEASER_CONFIG}

.PHONY: goreleaser-dev-release
goreleaser-dev-release: ## run goreleaser snapshot release
	./tools/bin/goreleaser_wrapper release --snapshot --rm-dist --config ${GORELEASER_CONFIG}

help:
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
