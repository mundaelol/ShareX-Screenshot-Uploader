FROM golang:1.23.0-alpine AS base

WORKDIR /app

FROM base AS builder

RUN apk add --no-cache git

RUN git clone https://github.com/mundaelol/ShareX-Screenshot-Uploader.git .
RUN go mod download && go mod verify
RUN go build -o /app/uploader main.go

FROM base AS runner

COPY --from=builder /app/.env /app/.env
COPY --from=builder /app/uploader /app/uploader

RUN chown -R nobody:nogroup /app

USER nobody

EXPOSE 8080/tcp

CMD ["/app/uploader"]
