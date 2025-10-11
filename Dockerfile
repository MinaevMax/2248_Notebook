FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


WORKDIR /app/cmd/server

RUN go build -o /app/main .

WORKDIR /app
EXPOSE 8080
CMD ["./main"]