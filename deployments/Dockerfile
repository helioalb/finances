# Build stage
FROM golang:1.26.2-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:latest

ARG VERSION=latest
ARG BUILD_DATE=unknown
ARG GIT_COMMIT=unknown
LABEL org.opencontainers.image.version=$VERSION \
	org.opencontainers.image.created=$BUILD_DATE \
	org.opencontainers.image.revision=$GIT_COMMIT \
  org.opencontainers.image.title="finances" \
	org.opencontainers.image.description="API de financas em Go" \
	org.opencontainers.image.source="https://github.com/helioalb/finances" \
	org.opencontainers.image.url="https://github.com/helioalb/finances" \
	org.opencontainers.image.licenses="MIT" \
	org.opencontainers.image.authors="Helio Albano"

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
