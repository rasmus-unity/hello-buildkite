FROM golang:1.15 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY server/main.go server/

RUN CGO_ENABLED=0 go build -o server server/*


FROM alpine:3

WORKDIR /app

COPY --from=builder /app/server/main ./server

CMD ["./server"]
