FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o bin ./consumer/main.go


ENTRYPOINT ["/app/bin"]