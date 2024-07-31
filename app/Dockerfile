FROM golang:1.20.6-alpine

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /api
EXPOSE 8000

CMD ["/api"]