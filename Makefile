.DELETE_ON_ERROR: clean

EXECUTABLES = go
K := $(foreach exec,$(EXECUTABLES),\
  $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

######## Variables ########
APP_NAME 	        ?= $(shell go list . | cut -d '/' -f 3)
APP_PACKAGE       ?= $(shell go list .)
APP_DEPENDENCIES  := $(shell go list -m -f '{{if not (or .Indirect .Main)}}{{.Path}}{{end}}' all)
APP_COVERAGE_FILE ?= $(BUILD_DIR)/coverage.txt
APP_COVERAGE_MODE	?= atomic

GIT_VERSION     ?= $(shell git rev-parse --abbrev-ref HEAD | cut -d "/" -f 2)
GIT_COMMIT      ?= $(shell git rev-parse HEAD | tr -d '\040\011\012\015\n')
GIT_BRANCH      ?= $(shell git rev-parse --abbrev-ref HEAD | tr -d '\040\011\012\015\n')
GIT_USER        ?= $(shell git config --get user.name)
BUILD_DATE      ?= $(shell date +'%Y-%m-%dT%H:%M:%S')
BUILD_DIR       := ./build
DIST_DIR        := ./dist
DIST_ASSEST_DIR := $(DIST_DIR)/assets

GO_CGO_ENABLED ?= 0
GO_OPTS        ?= -v
GO_OS          ?= darwin linux windows
GO_ARCH        ?= arm64 amd64
GO_HOST_OS     ?= $(shell $(GO) env GOHOSTOS)
GO_HOST_ARCH   ?= $(shell $(GO) env GOHOSTARCH)
GO_FILES       := $(shell go list ./... | grep -v /mocks/)
GO_GRAPH_FILE  := $(BUILD_DIR)/go-mod-graph.txt

GO_LDFLAGS_OPTIONS ?= -s -w
define EXTRA_GO_LDFLAGS_OPTIONS
-X '"'$(APP_PACKAGE)/internal/version.Version=$(GIT_VERSION)'"' \
-X '"'$(APP_PACKAGE)/internal/version.GitCommit=$(GIT_COMMIT)'"' \
-X '"'$(APP_PACKAGE)/internal/version.GitBranch=$(GIT_BRANCH)'"' \
-X '"'$(APP_PACKAGE)/internal/version.BuildUser=$(GIT_USER)'"' \
-X '"'$(APP_PACKAGE)/internal/version.BuildDate=$(BUILD_DATE)'"'
endef

GO_LDFLAGS := -ldflags "$(GO_LDFLAGS_OPTIONS) $(EXTRA_GO_LDFLAGS_OPTIONS)"

# compile all when cgo is disabled (slow)
ifeq ($(GO_CGO_ENABLED),0)
	GO_OPTS := $(GO_OPTS) -a
endif

######## Functions ########
# this is a funtion that will execute a command and print a message
# MAKE_DEBUG=true make venv-dev will print the command
# MAKE_STOP_ON_ERRORS=true make any fail will stop the execution if the command fails, this is useful for CI
# NOTE: if the dommand has a > it will print the output into the original redirect of the command
define exec_cmd
$(if $(filter $(MAKE_DEBUG),true),\
	$1 \
, \
	$(if $(filter $(MAKE_STOP_ON_ERRORS),true),\
		@$1  > /dev/null 2>&1 && printf "  🤞 ${1} ✅\n" || (printf "  ${1} ❌ 🖕\n"; exit 1) \
	, \
		$(if $(findstring >, $1),\
			@$1 2>/dev/null && printf "  🤞 ${1} ✅\n" || printf "  ${1} ❌ 🖕\n" \
		, \
			@$1 > /dev/null 2>&1 && printf '  🤞 ${1} ✅\n' || printf '  ${1} ❌ 🖕\n' \
		) \
	) \
)

endef # don't remove the whiteline before endef


######## Targets ########
##@ all
.PHONY: all
all: help

##@ build
.PHONY: build
build: go-generate go-fmt go-vet test ## Build the application
	@printf "👉 Building application...\n"
	$(call exec_cmd, CGO_ENABLED=$(GO_CGO_ENABLED) go build $(GO_LDFLAGS) $(GO_OPTS) -o $(BUILD_DIR)/$(APP_NAME) . )
	$(call exec_cmd, chmod +x $(BUILD_DIR)/$(APP_NAME) )

##@ build-dist
build-dist: ## Build the application for all platforms defined in GO_OS and GO_ARCH in this Makefile
	@printf "👉 Building application for different platforms...\n"
	$(foreach GOOS, $(GO_OS), \
		$(foreach GOARCH, $(GO_ARCH), \
			$(foreach app_name, $(APP_NAME), \
				$(call exec_cmd, GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(GO_CGO_ENABLED) go build $(GO_LDFLAGS) $(GO_OPTS) -o $(DIST_DIR)/$(app_name)-$(GOOS)-$(GOARCH) .) \
				$(call exec_cmd, chmod +x $(DIST_DIR)/$(app_name)-$(GOOS)-$(GOARCH)) \
			)\
		)\
	)

# this is needed to create the dist folder and the coverage file
$(APP_COVERAGE_FILE):
	@printf "👉 Creating coverage file...\n"
	$(call exec_cmd, mkdir -p $(BUILD_DIR) )
	$(call exec_cmd, touch $(APP_COVERAGE_FILE) )

##@ test
.PHONY: test
test: $(APP_COVERAGE_FILE) go-mod-tidy go-generate ## Run tests
	@printf "👉 Running tests...\n"
	$(call exec_cmd, go test \
										-v -race \
										-coverprofile=$(APP_COVERAGE_FILE) \
										-covermode=$(APP_COVERAGE_MODE) \
										./... \
	)

##@ lint
.PHONY: lint
lint: ## Run linters
	@printf "👉 Running linters...\n"
	$(call exec_cmd, golangci-lint run --timeout 5m)

##@ go-fmt
.PHONY: go-fmt
go-fmt: ## Format go code
	@printf "👉 Formatting go code...\n"
	$(call exec_cmd, go fmt ./... )

##@ go-vet
.PHONY: go-vet
go-vet: ## Vet go code
	@printf "👉 Vet go code...\n"
	$(call exec_cmd, go vet ./... )


##@ go-genereate
.PHONY: go-generate
go-generate: ## Generate go code
	@printf "👉 Generating go code...\n"
	$(call exec_cmd, go generate ./... )

##@ go-mod-tidy
.PHONY: go-mod-tidy
go-mod-tidy: ## Clean go.mod and go.sum
	@printf "👉 Cleaning go.mod and go.sum...\n"
	$(call exec_cmd, go mod tidy)

##@ go-mod-update
.PHONY: go-mod-update
go-mod-update: go-mod-tidy ## Update go.mod and go.sum
	@printf "👉 Updating go.mod and go.sum...\n"
	$(foreach DEP, $(APP_DEPENDENCIES), \
		$(call exec_cmd, go get -u $(DEP)) \
	)

##@ go-mod-vendor
.PHONY: go-mod-vendor
go-mod-vendor: ## Create mod vendor
	@printf "👉 Creating mod vendor...\n"
	$(call exec_cmd, go mod vendor)

##@ go-mod-verify
.PHONY: go-mod-verify
go-mod-verify: ## Verify go.mod and go.sum
	@printf "👉 Verifying go.mod and go.sum...\n"
	$(call exec_cmd, go mod verify)

##@ go-mod-download
.PHONY: go-mod-download
go-mod-download: ## Download go dependencies
	@printf "👉 Downloading go dependencies...\n"
	$(call exec_cmd, go mod download)

##@ go-mod-graph
.PHONY: go-mod-graph
go-mod-graph: ## Create a file with the go dependencies graph in build dir
	@printf "👉 Printing go dependencies graph...\n"
	$(call exec_cmd, go mod graph > $(GO_GRAPH_FILE))

##@ clean
.PHONY: clean
clean: ## Clean the environment
	@printf "👉 Cleaning environment...\n"
	$(call exec_cmd, go clean -n -x -i)
	$(call exec_cmd, rm -rf ./*.out)
	$(call exec_cmd, rm -rf ./logs)
	$(call exec_cmd, rm -rf $(BUILD_DIR))
	$(call exec_cmd, rm -rf $(DIST_DIR))

##@ help
.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##";                                             \
		printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ \
		{ printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/            \
		{ printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } '                  \
		$(MAKEFILE_LIST)