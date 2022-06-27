FROM golang:1.18 AS builder

WORKDIR /go/src/workdir
COPY go.* ./
RUN go mod download
COPY *.go .
RUN go mod tidy
RUN go test ./...
RUN GOOS=linux CGO_ENABLED=0 go build -o service .

FROM alpine:3.16
WORKDIR /root
COPY --from=builder /go/src/workdir/service .

EXPOSE 80
ENV NAME=""

ENTRYPOINT ["./service"]