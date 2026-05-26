run:
	mkdir -p bin
	go build -o bin/aerogo ./cmd/main && ./bin/aerogo

producer:
	go run ./cmd/producer

consumer:
	go run ./cmd/consumer

clean:
	go mod tidy && go mod download
