FROM golang:1.22-alpine3.18 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .
RUN go mod download

RUN go build -o main ./*.go

FROM scratch

COPY --from=builder /build/main /

EXPOSE 8080

ENTRYPOINT ["/main"]