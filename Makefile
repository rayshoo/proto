.DEFAULT_GOAL := all
GOOGLE_APIS = api
GATEWAY_PROTOS = gsm

pre-set:
	echo '' >> $${HOME}/.bashrc
	echo 'export PATH="$$PATH:$$(go env GOPATH)/bin"' >> $${HOME}/.bashrc
	set +e; mkdir $${HOME}/bin; set -e;
	echo 'export PATH="$$PATH:$${HOME}/bin"' >> $${HOME}/.bashrc
	curl --create-dirs -LO --output-dir $${HOME}/tmp \
	https://github.com/protocolbuffers/protobuf/releases/download/v3.15.0/protoc-3.15.0-win64.zip
	curl --create-dirs -LO --output-dir $${HOME}/tmp \
	https://github.com/grpc/grpc-web/releases/download/1.5.0/protoc-gen-grpc-web-1.5.0-windows-x86_64.exe
	curl --create-dirs -LO --output-dir $${HOME}/tmp \
	https://github.com/protocolbuffers/protobuf-javascript/releases/download/v3.21.2/protobuf-javascript-3.21.2-win64.zip
	curl --create-dirs -LO --output-dir $${HOME}/tmp \
	https://github.com/googleapis/googleapis/archive/refs/heads/master.zip
	mv $${HOME}/tmp/master.zip $${HOME}/tmp/googleapis.zip
.PHONY:pre-set

set:
	rm -rf $${HOME}/tmp/protoc && mkdir $${HOME}/tmp/protoc && \
	unzip $${HOME}/tmp/protoc-3.15.0-win64.zip -d $${HOME}/tmp/protoc && \
	if [ -d "$${HOME}/tmp/protoc/include" ] && [ ! -d "$${HOME}/include" ]; then mv $${HOME}/tmp/protoc/include $${HOME}/include; fi
	if [ -f "$${HOME}/tmp/protoc/bin/protoc.exe" ]; then mv $${HOME}/tmp/protoc/bin/protoc.exe $${HOME}/bin; fi

	if [ -f "$${HOME}/tmp/protoc-gen-grpc-web-1.5.0-windows-x86_64.exe" ]; then mv $${HOME}/tmp/protoc-gen-grpc-web-1.5.0-windows-x86_64.exe $${HOME}/bin/protoc-gen-grpc-web.exe; fi

	rm -rf $${HOME}/tmp/protobuf-javascript && mkdir $${HOME}/tmp/protobuf-javascript && \
	unzip $${HOME}/tmp/protobuf-javascript-3.21.2-win64.zip -d $${HOME}/tmp/protobuf-javascript && \
	if [ -f "$${HOME}/tmp/protobuf-javascript/bin/protoc-gen-js.exe" ] && [ ! -f "$${HOME}/bin/protoc-gen-js.exe" ]; then mv $${HOME}/tmp/protobuf-javascript/bin/protoc-gen-js.exe $${HOME}/bin; fi

	if [ ! -d "$${HOME}/tmp/googleapis/googleapis-master" ]; then unzip $${HOME}/tmp/googleapis.zip -d $${HOME}/tmp/googleapis; fi
	@for GOOGLE_API in $(GOOGLE_APIS); \
	do \
		if [ ! -d "$${HOME}/include/google/$$GOOGLE_API" ]; then mv $${HOME}/tmp/googleapis/googleapis-master/google/$$GOOGLE_API $${HOME}/include/google; fi \
	done

	go install golang.org/x/tools/cmd/goimports@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

	go install github.com/cosmtrek/air@v1.52.1
.PHONY:set

imports:
	goimports -l -w .
.PHONY:imports

proto:
	protoc --go_out=. --go-grpc_out=. *.proto
	@for GATEWAY_PROTO in $(GATEWAY_PROTOS); \
	do \
		protoc --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt generate_unbound_methods=true $${GATEWAY_PROTO}.proto; \
	done
.PHONY:proto

all: proto imports
.PHONY:all