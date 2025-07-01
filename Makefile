.PHONY: test_release run_docker update_depends update_frontend_depends build

EXT :=
MOVE := mv
RM_RF := rm -rf ./static

UNAME_S := $(shell uname -s)

ifeq ($(OS),Windows_NT)
	EXT := .exe
	MOVE := move
	RM_RF := if exist static ( rmdir /s /q static )
endif

test_release:
	@echo ">>> Starting test release"
	@goreleaser release --snapshot --clean
	@echo ">>> Test release complete"

run_docker:
	@echo ">>> Starting Docker containers"
	@docker-compose up -d
	@echo ">>> Docker containers are up"

update_depends:
	@echo ">>> Updating Go dependencies"
	@go get all
	@go mod tidy
	@echo ">>> Go dependencies updated"

update_frontend_depends:
	@echo ">>> Updating frontend dependencies"
	@cd Frontend && bun update
	@echo ">>> Frontend dependencies updated"

build:
	@echo ">>> Building Frontend"
	@(cd Frontend && bun install && bun run build)
	@echo ">>> Cleaning static folder"
	@$(RM_RF)
	@echo ">>> Moving new static files"
	@$(MOVE) Frontend/static .
	@echo ">>> Building Go Backend"
	@go build -mod=readonly -trimpath \
		-tags="mysql postgres sqlite sqlserver" \
		-o=QuicKNote$(EXT) \
		-ldflags="-s -w -buildid= -extldflags=-static" \
		-gcflags="all=-d=ssa/check_bce/debug=0" \
		-asmflags="-trimpath" main.go
	@echo ">>> Build complete: ./QuicKNote$(EXT)"
