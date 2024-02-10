FROM --platform=linux/amd64 golang:1.21.6-bullseye as dependencies
WORKDIR /app
COPY ../server-side/go.mod go.sum ./
RUN go mod download

FROM --platform=linux/amd64 golang:1.21.6-bullseye as builder
WORKDIR /app
COPY --from=dependencies /go/pkg /go/pkg
COPY ../server-side ./
RUN GOOS=linux go build -v -o services-server

FROM --platform=linux/amd64 debian:bullseye-20240110-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*
ENV ENVIRONMENT=production
COPY --from=builder /app/api-server /api-server
CMD ["/api-server"]
