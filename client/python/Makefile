PYTHOB_PROTO_SRC=pb

.PHONY: init
# init env
init:

	pip install grpcio-tools

.PHONY: pb-meta
pb-meta:
	python -m grpc_tools.protoc -I./pb \
		--python_out=./pb/grpc \
		--pyi_out=./pb/grpc \
		--grpc_python_out=./pb/grpc ./pb/meta.proto

.PHONY: pb-user
pb-user:
	python -m grpc_tools.protoc -I./pb \
		--python_out=./pb/grpc \
		--pyi_out=./pb/grpc \
		--grpc_python_out=./pb/grpc ./pb/user.proto

.PHONY: pb-service
pb-service:
	python -m grpc_tools.protoc -I./pb \
		--python_out=./pb/grpc \
		--pyi_out=./pb/grpc \
		--grpc_python_out=./pb/grpc ./pb/service.proto

.PHONY: pb-empty
pb-empty:
	python -m grpc_tools.protoc -I./pb \
		--python_out=./pb/grpc \
		--pyi_out=./pb/grpc ./pb/empty.proto

.PHONY: pb-annotations
pb-anno:
	python -m grpc_tools.protoc -I./pb \
		--python_out=./pb/grpc \
		--pyi_out=./pb/grpc ./pb/annotations.proto

.PHONY: pb-http
pb-http:
	python -m grpc_tools.protoc -I./pb \
		--python_out=./pb/grpc \
		--pyi_out=./pb/grpc ./pb/http.proto

.PHONY: pb-timestamp
pb-ts:
	python -m grpc_tools.protoc -I./pb \
		--python_out=./pb/grpc \
		--pyi_out=./pb/grpc ./pb/timestamp.proto

