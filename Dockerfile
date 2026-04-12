# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS build
WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/todo-api ./cmd/server

FROM alpine:3.20
RUN apk add --no-cache ca-certificates \
	&& adduser -D -H -u 65532 appuser

USER appuser:appuser
COPY --from=build /out/todo-api /todo-api

EXPOSE 8080
ENV PORT=8080

ENTRYPOINT ["/todo-api"]
