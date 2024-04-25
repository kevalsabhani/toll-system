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
	@./bin/distance_calculator

invoice:
	@go build -o bin/invoice_generator invoice_generator/main.go
	@./bin/invoice_generator

.PHONY: obuclient obuserver kafka distance_cal invoice