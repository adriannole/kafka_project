FROM golang:1.20
WORKDIR /app
COPY producer/producer.go .
RUN go mod init producer && go mod tidy
RUN go build -o producer .
EXPOSE 8080
CMD ["./producer"]
