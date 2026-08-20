# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Go & Tools
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
GO_COMPILER        = go
GO_LINTER          = go
GO_STATIC_ANALYZER = go
DOCKER_BIN         = docker

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Application
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
APP_NAME    = kneesocks
APP_VERSION = v1.0.2
APP_RELEASE = prod
APP_ROOT    = .

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Application - Sources
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
SRC_DIR = $(APP_ROOT)/internal
BIN_DIR = $(APP_ROOT)/cmd

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Application - Build
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
BUILD_DIR        = $(APP_ROOT)/build
BUILD_ENTRYPOINT = $(APP_ROOT)/$(BIN_DIR)/$(APP_NAME)/main.go
BUILD_TARGET     = $(APP_ROOT)/$(BUILD_DIR)/$(APP_VERSION)/$(APP_NAME)


# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Application - Deploy
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
DEPLOY_DIR                      = $(APP_ROOT)/deploy
DEPLOY_PROD_DIR                 = $(DEPLOY_DIR)/prod
DEPLOY_DEV_DIR                  = $(DEPLOY_DIR)/dev
DEPLOY_TEST_DIR                 = $(DEPLOY_DIR)/test

DEPLOYED_DOCKERIGNORE_FILE      = $(APP_ROOT)/.dockerignore
DEPLOYED_COMPOSE_FILE           = $(APP_ROOT)/compose.yml
DEPLOYED_AIR_FILE               = $(APP_ROOT)/.air.toml
DEPLOYED_APP_DOCKERFILE         = $(APP_ROOT)/app.Dockerfile
DEPLOYED_SERVICE_DOCKERFILE     = $(APP_ROOT)/test_server.Dockerfile
DEPLOYED_E2E_DOCKERFILE         = $(APP_ROOT)/e2e.Dockerfile

DEPLOYED_APP_CONTAINER_NAME     = kneesocks-app
DEPLOYED_APP_IMAGE_NAME         = kneesocks-app:latest
DEPLOYED_SERVICE_CONTAINER_NAME = kneesocks-test-service
DEPLOYED_SERVICE_IMAGE_NAME     = kneesocks-test-service:latest
DEPLOYED_E2E_CONTAINER_NAME     = kneesocks-e2e
DEPLOYED_E2E_IMAGE_NAME         = kneesocks-e2e:latest

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Application - Testing
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
E2E_NAME           = e2e
E2E_TESTS_DIR      = $(APP_ROOT)/test/$(E2E_NAME)
E2E_BINARY_TARGET  = $(BUILD_DIR)/$(APP_VERSION)/$(E2E_NAME)
UNIT_TEST_PACKAGES = $(SRC_DIR) $(BIN_DIR) ./pkg/...
SERVICE_NAME       = test_server

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# PHONY & such
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
PHONY_TARGETS  = all
PHONY_TARGETS += all-tests all-build all-code
PHONY_TARGETS += run-compose
PHONY_TARGETS += run-tests-e2e-compose run-tests-unit-compose
PHONY_TARGETS += tests-e2e-compose tests-unit-compose
PHONY_TARGETS += application-compose
PHONY_TARGETS += deploy-prod deploy-test
PHONY_TARGETS += code-inspect code-style
PHONY_TARGETS += clean
PHONY_TARGETS += help

.PHONY: $(PHONY_TARGETS)
.DEFAULT_GOAL := all

all: all-tests all-build all-code

all-tests: tests-e2e-compose tests-unit-compose

all-build: application-compose-images

all-code: code-inspect code-style

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Run
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
run-compose: application-compose-images
	$(DOCKER_BIN) compose up

run-tests-e2e-compose: application-compose-images
	$(DOCKER_BIN) compose up

run-tests-unit-compose:

tests-e2e-compose: deploy-$(APP_RELEASE)

tests-unit-compose: deploy-$(APP_RELEASE)

application-compose-images: deploy-$(APP_RELEASE)
	$(DOCKER_BIN) compose build

deploy-prod:
	cp $(DEPLOY_PROD_DIR)/.dockerignore $(DEPLOYED_DOCKERIGNORE_FILE)
	cp $(DEPLOY_PROD_DIR)/compose.yml $(DEPLOYED_COMPOSE_FILE)
	cp $(DEPLOY_PROD_DIR)/Dockerfile $(DEPLOYED_APP_DOCKERFILE)

deploy-dev:
	cp $(DEPLOY_DEV_DIR)/compose.yml $(DEPLOYED_COMPOSE_FILE)
	cp $(DEPLOY_DEV_DIR)/.air.toml $(DEPLOYED_AIR_FILE)

deploy-test:
	cp $(DEPLOY_TEST_DIR)/compose.yml $(DEPLOYED_COMPOSE_FILE)
	cp $(DEPLOY_TEST_DIR)/app.Dockerfile $(DEPLOYED_APP_DOCKERFILE)
	cp $(DEPLOY_TEST_DIR)/test_server.Dockerfile $(DEPLOYED_SERVICE_DOCKERFILE)
	cp $(DEPLOY_TEST_DIR)/e2e.Dockerfile $(DEPLOYED_E2E_DOCKERFILE)

code-inspect:
	$(GO_STATIC_ANALYZER) vet ./...

code-style:
	$(GO_LINTER) fmt ./...

clean-deployment: clean-deployment-$(APP_RELEASE)

clean-deployment-prod:
	rm -rf $(BUILD_DIR)
	rm -f $(DEPLOYED_DOCKERIGNORE_FILE)
	rm -f $(DEPLOYED_COMPOSE_FILE)
	rm -f $(DEPLOYED_APP_DOCKERFILE)

clean-deployment-dev:
	rm -rf $(BUILD_DIR)
	rm -f $(DEPLOYED_COMPOSE_FILE)
	rm -f $(DEPLOYED_AIR_FILE)

clean-deployment-test:
	rm -rf $(BUILD_DIR)
	rm -f $(DEPLOYED_COMPOSE_FILE)
	rm -f $(DEPLOYED_APP_DOCKERFILE)
	rm -f $(DEPLOYED_SERVICE_DOCKERFILE)
	rm -f $(DEPLOYED_E2E_DOCKERFILE)

help:
	@echo "Make scripts for track-my-tasks app:"
	@echo ""
	@echo "  all                     tests, build and code-quality checks (default)"
	@echo "  all-tests               build the e2e binary and run unit tests"
	@echo "  all-build               build the app binary, compose files and docker image"
	@echo "  all-code                run static checks and format code"
	@echo ""
	@echo "  run-compose             copy deploy/$(APP_RELEASE) files to root and docker compose up"
	@echo ""
	@echo "  run-tests-unit-compose  run unit tests"
	@echo "  run-tests-e2e-compose   copy deploy/test files to root and docker compose up"
	@echo ""
	@echo "  code-inspect            run static analysis ($(GO_STATIC_ANALYZER) vet)"
	@echo "  code-style              format the codebase ($(GO_LINTER) fmt)"
	@echo ""
	@echo "  clean                   remove build artifacts and copied deploy files"
	@echo ""
	@echo "Configuration (override with VAR=value, e.g. make all GO_COMPILER=go1.22):"
	@echo "  GO_COMPILER             go binary to build/test with (default: $(GO_COMPILER))"
	@echo "  GO_LINTER               go binary used by code-style (default: $(GO_LINTER))"
	@echo "  GO_STATIC_ANALYZER      go binary used by code-inspect (default: $(GO_STATIC_ANALYZER))"
	@echo "  DOCKER_BIN              docker binary/path (default: $(DOCKER_BIN))"
	@echo "  APP_NAME                binary/image name (default: $(APP_NAME))"
	@echo "  APP_VERSION             build/image version (default: $(APP_VERSION))"
	@echo "  APP_RELEASE             prod or test - picks deploy/ config for application-compose (default: $(APP_RELEASE))"
	@echo "  APP_ROOT                project root, rarely needs changing (default: $(APP_ROOT))"