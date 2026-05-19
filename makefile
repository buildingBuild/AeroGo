run:
	mkdir -p bin
	go build -o bin/aerogo ./cmd/main && ./bin/aerogo
