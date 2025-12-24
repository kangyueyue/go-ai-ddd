.PHONY: run_frontend, build_frontend, run, tidy

run_frontend:
	cd vue-frontend && npm run serve
	@echo "run frontend success"

build_frontend:
	cd vue-frontend && npm run build
	@echo "build frontend success"

run:
	go run cmd/main.go
	@echo "go run success"

tidy:
	go mod tidy
	@echo "go tidy success"