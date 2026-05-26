run:
	mkdir -p bin
	go build -o bin/aerogo ./cmd/main && ./bin/aerogo

clean:
	go mod tidy && go mod download
