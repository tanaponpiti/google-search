FROM --platform=linux/amd64 golang:1.21.6-bullseye as base
WORKDIR /app

FROM base AS dependencies
COPY ./go.mod go.sum ./
RUN go mod download

FROM dependencies AS builder
WORKDIR /app
COPY --from=dependencies /go/pkg /go/pkg
COPY . ./
RUN GOOS=linux go build -v -o api-server

FROM --platform=linux/amd64 debian:bullseye-20240110-slim AS serve-stage
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*
ENV GIN_MODE=release
COPY --from=builder /app/api-server /api-server
CMD ["/api-server"]
