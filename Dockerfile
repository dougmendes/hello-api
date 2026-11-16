FROM golang:1.27 as builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ARG VERSION

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags="-s -w -X main.version=$VERSION" -o hello-api .

FROM alpine:3.24

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

RUN addgroup -S appuser \
    && adduser -S -G appuser -H -s /sbin/nologin appuser

COPY --from=builder --chown=appuser:appuser /app/hello-api /app/hello-api

USER appuser

ENTRYPOINT ["/app/hello-api"]
