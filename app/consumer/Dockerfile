FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o bin /app/main.go

ENTRYPOINT ["/app/bin"]