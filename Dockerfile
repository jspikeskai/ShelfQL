FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY app/ .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ShelfQL main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/ShelfQL .
CMD ["./ShelfQL"]