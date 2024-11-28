FROM golang:1.23.0-alpine

WORKDIR /app

RUN apk add --no-cache git && git clone https://github.com/mundaelol/ShareX-Screenshot-Uploader.git .

RUN go mod download && go mod verify

EXPOSE 8080

CMD ["go", "run", "main.go"]
