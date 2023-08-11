FROM golang:1.18.4-alpine as builder

RUN apk update && apk add --no-cache bash protobuf-dev
RUN apk add

RUN mkdir /build
WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/main ./cmd

FROM alpine:latest
COPY --from=builder /build/bin /bin/
COPY --from=builder /build/config/config.yaml /config/
COPY --from=builder /build/.env .

ENTRYPOINT ["/bin/main"]