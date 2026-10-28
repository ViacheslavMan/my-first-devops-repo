FROM golang:1.22-alpine
WORKDIR /app
COPY main.go .
RUN go mod init app && go mod tidy && go build -o server .
CMD ["./server"]
