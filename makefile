NPX = pnpm exec
PROTO = $(wildcard proto/*.proto)

ifeq ($(OS),Windows_NT)
	EXT = .exe
endif

TARGET_TRIPLE = $(shell rustc -vV | sed -n 's|host: ||p')

server-bin:
	cd server && go build -o server-$(TARGET_TRIPLE)$(EXT)

server/api:
	mkdir server/api

src/proto:
	mkdir src/proto

protobuf: server/api src/proto
	protoc -I=proto --go_out=server/ --go-grpc_out=server/ $(PROTO)
	$(NPX) protoc --ts_out src/proto \
		--ts_opt long_type_string \
		--ts_opt optimize_code_size \
		--proto_path proto \
		$(PROTO)
