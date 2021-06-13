FROM golang:1.16-alpine as builder
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY opa-nginx-external-auth/ opa-nginx-external-auth/
RUN go build -o opa-nginx-external-auth echo/main.go

FROM alpine:3.13
RUN apk add --no-cache ca-certificates tini
COPY --from=builder /workspace/opa-nginx-external-auth /usr/local/bin/
WORKDIR /workspace
ENTRYPOINT [ "/sbin/tini", "--", "opa-nginx-external-auth" ]