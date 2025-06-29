.PHONY: test_release run_docker update_depends update_frontend_depends build

EXT :=
MOVE := mv
RM_RF := rm -rf ./static

UNAME_S := $(shell uname -s)

ifeq ($(OS),Windows_NT)
	EXT := .exe
	MOVE := move
	RM_RF := if exist static ( rmdir /s /q static )
# Darwin (macOS) already uses mv, so no special case needed.
endif

test_release:
	goreleaser release --snapshot --clean

run_docker:
	docker-compose up -d

update_depends:
	go get all && go mod tidy

update_frontend_depends:
	cd Frontend && bun update

build:
	@echo "--- Building Frontend ---"
	@(cd Frontend && bun install && bun run build)
	@echo "--- Cleaning old static files and moving new ones ---"
	$(RM_RF)
	$(MOVE) Frontend/static .
	@echo "--- Building Go Backend ---"
	go build -mod=readonly -trimpath \
		-tags="mysql postgres sqlite sqlserver" \
		-o=QuicKNote$(EXT) \
		-ldflags="-s -w -buildid= -extldflags=-static" \
		-gcflags="all=-d=ssa/check_bce/debug=0" \
		-asmflags="-trimpath" main.go
	@echo "--- Build complete: ./QuicKNote$(EXT) ---"