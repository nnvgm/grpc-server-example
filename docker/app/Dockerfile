FROM nnvgm/grpc-common-example as builder

WORKDIR $GOPATH/src/github.com/nnvgm/grpc-server-example

COPY . .

RUN go get -d -v ./... && go build -o app

FROM alpine:3.10

COPY --from=builder /go/src/github.com/nnvgm/grpc-server-example/app /usr/bin/app

EXPOSE 50001

CMD ["app"]