run:
	@go build -o bin/app cmd/app/main.go
	@./bin/app
migration-create:
	migrate create -ext json -dir ./migrations -seq $(filter-out $@,$(MAKECMDGOALS))