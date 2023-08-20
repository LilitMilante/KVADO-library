.PHONY: test up gen

test:
	go test -v -race ./...

up-docker:
	docker-compose up -d --build

up:
	docker run --rm --name library-db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=dev -e MYSQL_DATABASE=library -d mysql:8.1.0

down:
	docker rm -f library-db

gen:
	protoc --go_out=./gen --go_opt=paths=source_relative \
        --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
        ./proto/library.proto

deps:
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

send:
	grpcurl -insecure -d '{"author_id": "44a54d7b-6289-4b12-b030-1ffd884763cb"}' localhost:8080 proto.Library/BooksByAuthorID
