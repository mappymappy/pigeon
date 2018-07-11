gen-pb:
	cd proto && protoc --go_out=plugins=grpc:../pb chat_service.proto
build-server:
	go build -o ./pigeon_server
run-server: build-server
	./pigeon_server
build-client:
	go build -o ./pigeon_client ./client
run-client: build-client
	./pigeon_client
build-client%:
	go build -o ./pigeon_client$* ./client
run-client%: build-client%
	./pigeon_client$*
