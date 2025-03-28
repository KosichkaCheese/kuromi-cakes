FROM golang:1.24.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init

EXPOSE 8000

CMD swag init && go run main.go