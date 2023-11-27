gps-unit:
	@go build -o bin/gps-unit ./gps-unit/*.go
	@./bin/gps-unit

data-receiver:
	@go build -o bin/data_receiver ./data_receiver/*.go
	@./bin/data_receiver


.PHONY: gps-unit
