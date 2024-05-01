kafka:
	@docker compose down
	@docker compose up -d

obuclient:
	@go build -o bin/obu_data_sender client/main.go
	@./bin/obu_data_sender

obuserver:
	@go build -o bin/obu_data_receiver server/main.go
	@./bin/obu_data_receiver

distance_cal:
	@go build -o bin/distance_calculator distance_calculator/main.go
	@./bin/distance_calculator -t grpc

invoice:
	@go build -o bin/invoice_generator invoice_generator/main.go
	@./bin/invoice_generator -t grpc

proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pb/*.proto

.PHONY: obuclient obuserver kafka distance_cal invoice