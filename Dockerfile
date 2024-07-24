FROM golang:1.20.6-alpine

WORKDIR /app

COPY . .
RUN go mod download

EXPOSE 8000

CMD ["go", "run", "main.go"]