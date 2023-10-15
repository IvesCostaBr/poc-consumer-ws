FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o bin /app/consumer/main.go

ENTRYPOINT ["/app/bin"]