FROM golang:1.23.0 AS builder

WORKDIR /book

COPY . .

RUN go build -o app cmd/book/app.go

FROM golang:1.23.0

WORKDIR /book

COPY --from=builder /book/app .

COPY --from=builder /book/cmd/tls/localhost.pem /localhost.pem
COPY --from=builder /book/cmd/tls/localhost-key.pem /localhost-key.pem

CMD ["./app"]

