GOPATH ?=	$(HOME)/go
BIN :=		$(GOPATH)/bin/recettator
SOURCE :=	$(shell find . -name "*.go")
OWN_PACKAGES := $(shell go list ./... | grep -v vendor)


build: $(BIN)


$(BIN): $(SOURCE)
	go install ./cmd/recettator


.PHONY: docker
docker:
	docker build -t ultreme/recettator .


.PHONY: test
test:
	go test -v $(OWN_PACKAGES)
