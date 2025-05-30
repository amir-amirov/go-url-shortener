FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

# This is purely informational and not required for the build
# Making it easier to understand the image
EXPOSE 8080

CMD ["./main"]
