FROM golang:alpine3.11 AS builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/darkraiden/aws-kms-encrypter

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o aws-kms-encrypter ./main.go

FROM alpine:3.11.3

RUN adduser -S -D -H -h /app -u 1001 aws-kms-encrypter

USER aws-kms-encrypter

COPY --from=builder /go/src/github.com/darkraiden/aws-kms-encrypter/aws-kms-encrypter /app/aws-kms-encrypter

WORKDIR /app

ENTRYPOINT ["./aws-kms-encrypter"]
