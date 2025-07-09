FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o goshell main.go

CMD ["./goshell"]
