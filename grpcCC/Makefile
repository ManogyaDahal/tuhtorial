old_add:
	protoc \
		-I=server/proto \
		--go_out=server/proto/gen \
		--go_opt=paths=source_relative \
		server/proto/Add/add.proto

add:
	protoc \
		-I=server/proto \
		--go_out=server/proto/gen \
		--go_opt=paths=source_relative \
		--go-grpc_out=server/proto/gen \
		--go-grpc_opt=paths=source_relative \
		server/proto/Add/add.proto
