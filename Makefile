proto:
	sh third_party/protoc_gen.sh

build: proto
	go build -o server cmd/server/main.go

docker: proto 
	docker build  -t cardenasrjl/emem:latest .

build_client:
	go build -o client cmd/client-grpc/main.go

up: build
	docker-compose up -d
	
down:
	docker-compose down

restart: down up 

clean-up: down  
	docker system prune -f
	make docker
	make up