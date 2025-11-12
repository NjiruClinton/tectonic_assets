FROM golang:1.25-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod tidy

CMD ["air", "-d"]