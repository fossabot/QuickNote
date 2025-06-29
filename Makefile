.PHONY: test_release, run_docker

test_release:
	goreleaser release --snapshot --clean

run_docker:
	docker-compose up -d

update_depends:
	go get all && go mod tidy

update_frontend_depends:
	cd Frontend && bun update