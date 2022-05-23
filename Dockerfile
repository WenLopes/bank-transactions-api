FROM golang:1.18-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server ./api/main.go

EXPOSE 8080

ENTRYPOINT /app/server