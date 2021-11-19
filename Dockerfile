FROM golang:1.17 AS builder
WORKDIR /workspace

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download
COPY ./src .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM alpine:3.14
WORKDIR /
COPY --from=builder /main .

ENTRYPOINT ["/main"]
