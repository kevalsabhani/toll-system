obu_sender:
	@go build -o bin/obu_data_sender client/main.go
	@./bin/obu_data_sender

obu_receiver:
	@go build -o bin/obu_data_receiver server/main.go
	@./bin/obu_data_receiver


.PHONY: obu_sender obu_receiver