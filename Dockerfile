# syntax=docker/dockerfile:1

FROM golang:1.19

# Set destination for COPY
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN ls
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

EXPOSE 8080
CMD ["./main"]