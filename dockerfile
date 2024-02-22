FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o build/bin/go_codifin cmd/codifin/main.go
EXPOSE 8080
CMD ["./build/bin/go_codifin"]

