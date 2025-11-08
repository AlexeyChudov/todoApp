FROM golang:1.22-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o todoapp ./cmd/main/main.go

FROM alpine:3.20 AS runner
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/todoapp .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates
EXPOSE 3000
CMD ["./todoapp"]
