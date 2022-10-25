export GOBIN=$(PWD)/bin
export PROTOBUF_ROOT=$(PWD)/_vendor/protobuf-3.20.0

BUILTINS_SUPPORT_PKG_REL := support/types

.PHONY: install test gen-conformance gen-include genall

install:
	go install -tags protolegacy google.golang.org/protobuf/cmd/protoc-gen-go
	go install -tags protolegacy ./cmd/protoc-gen-go-vtproto
# 	go install -tags protolegacy github.com/gogo/protobuf/protoc-gen-gofast

go-generate:
	go generate ./internal/builtinprotos/descriptors

install-builtins-gen: go-generate
	go install -tags protolegacy ./internal/cmd/protoc-gen-go-vtproto-builtins

gen-builtins: install-builtins-gen
	$(PROTOBUF_ROOT)/src/protoc \
	  	--plugin protoc-gen-go-vtproto-builtins="${GOBIN}/protoc-gen-go-vtproto-builtins" \
		--go-vtproto-builtins_out="$(BUILTINS_SUPPORT_PKG_REL)" \
		--go-vtproto-builtins_opt=support_pkg_root="$(shell go list -m)/$(BUILTINS_SUPPORT_PKG_REL)" \
		--go-vtproto-builtins_opt=registry_file=./internal/builtinprotos/registry/registry.go \
		internal/builtinprotos/gen-helpers/dummy.proto

gen-prototype-helpers: install-builtins-gen
	$(PROTOBUF_ROOT)/src/protoc \
		--go-vtproto-builtins_out=support/types --plugin protoc-gen-go-vtproto-builtins="${GOBIN}/protoc-gen-go-vtproto-builtins" \
		--go-vtproto-builtins_opt=registry_file=./generator/builtins/builtins.go \
		--go-vtproto-builtins_opt=registry_package=builtins \
		dummy/dummy.proto

gen-conformance:
	$(PROTOBUF_ROOT)/src/protoc \
		--proto_path=$(PROTOBUF_ROOT) \
		--go_out=conformance --plugin protoc-gen-go="${GOBIN}/protoc-gen-go" \
		--go-vtproto_out=conformance --plugin protoc-gen-go-vtproto="${GOBIN}/protoc-gen-go-vtproto" \
		-I$(PROTOBUF_ROOT)/src \
		--go_opt=Msrc/google/protobuf/test_messages_proto2.proto=internal/conformance \
		--go_opt=Msrc/google/protobuf/test_messages_proto3.proto=internal/conformance \
		--go_opt=Mconformance/conformance.proto=internal/conformance \
		--go-vtproto_opt=Msrc/google/protobuf/test_messages_proto2.proto=internal/conformance \
		--go-vtproto_opt=Msrc/google/protobuf/test_messages_proto3.proto=internal/conformance \
		--go-vtproto_opt=Mconformance/conformance.proto=internal/conformance \
		src/google/protobuf/test_messages_proto2.proto \
		src/google/protobuf/test_messages_proto3.proto \
		conformance/conformance.proto

gen-include:
	$(PROTOBUF_ROOT)/src/protoc \
		--proto_path=include \
		--go_out=include --plugin protoc-gen-go="${GOBIN}/protoc-gen-go" \
		-I$(PROTOBUF_ROOT)/src \
		github.com/planetscale/vtprotobuf/vtproto/ext.proto
	mv include/github.com/planetscale/vtprotobuf/vtproto/*.go ./vtproto

gen-testproto:
	$(PROTOBUF_ROOT)/src/protoc \
		--proto_path=testproto \
		--proto_path=include \
		--go_out=. --plugin protoc-gen-go="${GOBIN}/protoc-gen-go" \
		--go-vtproto_out=allow-empty=true:. --plugin protoc-gen-go-vtproto="${GOBIN}/protoc-gen-go-vtproto" \
		-I$(PROTOBUF_ROOT)/src \
		testproto/empty/empty.proto \
		testproto/pool/pool.proto \
		testproto/pool/pool_with_slice_reuse.proto \
		testproto/proto3opt/opt.proto \
		testproto/proto2/scalars.proto \
		|| exit 1;

genall: install gen-include gen-conformance gen-testproto

test: install gen-conformance
	go test -short ./...
	go test -count=1 ./conformance/...
	GOGC="off" go test -count=1 ./testproto/pool/...
