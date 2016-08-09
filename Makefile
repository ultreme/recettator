BIN :=		recettator
SOURCE :=	$(shell find . -name "*.go")
OWN_PACKAGES := $(shell go list ./... | grep -v vendor)


build: $(BIN)


$(BIN): $(SOURCE)
	go build -o ./$@ ./cmd/$@/main.go


.PHONY: test
test:
	go test -v $(OWN_PACKAGES)
