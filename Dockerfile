FROM golang:1.11
COPY . /go/src/ultre.me/recettator
WORKDIR /go/src/ultre.me/recettator
RUN GO111MODULE=on go install -v ./cmd/recettator
ENTRYPOINT ["recettator"]
CMD ["-h"]
