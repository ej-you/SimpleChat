info_log = "/logs/info-log.log"
error_log = "/logs/error-log.log"

source = ./main.go
dest = ./go_app


dev:
	go run $(source)

lint:
	golangci-lint run -c ./.golangci.yml ./...

migrate:
	go run $(source) migrate

compile:
	go build -o $(dest) $(source)

prod:
	@echo "Running migrations..."
	/root/$(dest) migrate
	@echo "Running main app..."
	/root/$(dest) >> $(info_log) 2>> $(error_log)

swagger:
	@swag init
	@swag fmt
