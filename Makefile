kafka:
	@docker compose down
	@docker compose up -d

obuclient:
	@go build -o bin/obu_data_sender client/main.go
	@./bin/obu_data_sender

obuserver:
	@go build -o bin/obu_data_receiver server/main.go
	@./bin/obu_data_receiver

distance:
	@go build -o bin/distance_calculator distance_calculator/main.go
	@./bin/distance_calculator


.PHONY: obuclient obuserver kafka distance