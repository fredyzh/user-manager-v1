JAVA_META_PROTO_SRC=src/main/java/com/client/test/pb

.PHONY: java client meta

pb-java-meta:
	
	protoc -I=$(JAVA_META_PROTO_SRC)  \
           --java_out=$(JAVA_META_PROTO_SRC) \
           $(JAVA_META_PROTO_SRC)/meta.proto

.PHONY: java client user

pb-java-user:
	
	protoc -I=$(JAVA_META_PROTO_SRC) \
           --java_out=$(JAVA_META_PROTO_SRC)/ \
           $(JAVA_META_PROTO_SRC)/user.proto


pb-java-service:
	
	protoc --plugin=protoc-gen-grpc-java=C:\tools\grpc_java\protoc-gen-grpc-java-1.58.0-windows-x86_64.exe \
		   --proto_path=$(JAVA_META_PROTO_SRC) \
		   --java_out=$(JAVA_META_PROTO_SRC) \
		   --grpc-java_out=$(JAVA_META_PROTO_SRC)/ \
           $(JAVA_META_PROTO_SRC)/service.proto