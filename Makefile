kafka:
	@docker compose down
	@docker compose up -d

obuclient:
	@go build -o bin/obu_data_sender client/main.go
	@./bin/obu_data_sender

obuserver:
	@go build -o bin/obu_data_receiver server/main.go
	@./bin/obu_data_receiver


.PHONY: obuclient obuserver kafka