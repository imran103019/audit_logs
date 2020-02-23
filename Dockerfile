FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

LABEL maintainer="rajs.imran93@gmail.com"

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o main .
EXPOSE 8080

# Run the executable
CMD ["./main"]