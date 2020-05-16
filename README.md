# GRPC Server Example

```bash
# Start server
$ GRPC_PORT=50001 go run main.go

# build
$ go build -o app
# start
$ GRPC_PORT=50001 ./app

# build image
$ docker build -f ./docker/app/Dockerfile -t ${USERNAME}/grpc-server-example .
$ docker push ${USERNAME}/grpc-server-example
```
