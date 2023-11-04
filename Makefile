.PHONY: all  test

all: fmt vet  test

build:
	bash hack/build.sh

test: fmt vet
	go test -v ./pkg/... ./cmd/... -coverprofile cover.out

fmt:
	go fmt ./pkg/... ./cmd/...

vet:
	go vet ./pkg/... ./cmd/...

clean:
	-rm -Rf bin

codegen:
	bash hack/update-codegen.sh
