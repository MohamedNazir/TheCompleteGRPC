gen:
#	protoc --proto_path=proto --go_out=plugins=grpc:. proto/*.proto
	protoc --proto_path=proto --go_out=plugins=grpc:pb proto/*.proto

clean:
	rm -R pb/*.*

server:
	go run cmd/server/main.go -port 9090

client:
	go run cmd/client/main.go -address 0.0.0.0:9090

test:
	go test -cover -race ./...