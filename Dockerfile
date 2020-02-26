ARG GO_VERSION=1.14

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir /gin-template
WORKDIR /gin-template

COPY . .

RUN go mod download
RUN go mod verify

RUN go build


FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /gin-template
WORKDIR /gin-template
COPY --from=builder /gin-template .

EXPOSE 8080

ENTRYPOINT ["./gin-template"]

