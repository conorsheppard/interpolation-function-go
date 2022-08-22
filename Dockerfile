FROM golang:1.18-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
ENTRYPOINT exec go test -v ./...