FROM golang:1.16-alpine as builder
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY echo/ echo/
RUN go build -o httpecho echo/main.go

FROM alpine:3.13
RUN apk add --no-cache ca-certificates tini
COPY --from=builder /workspace/httpecho /usr/local/bin/
WORKDIR /workspace
ENTRYPOINT [ "/sbin/tini", "--", "httpecho" ]