proto:
	sh third_party/protoc_gen.sh

build: proto
	go build -o server cmd/server/main.go

build_client:

	go build -o client cmd/client-grpc/main.go

up: build
	docker-compose up -d
	./server

down:
	docker-compose down

restart: down up 