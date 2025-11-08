FROM golang:1.22-alpine AS builder

RUN apk add --no-cache bash build-base ca-certificates
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Важное для мультиарх сборки: статическая сборка без CGO
# BuildKit автоматически установит GOOS/GOARCH под текущую платформу сборки
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    GOFLAGS="-ldflags=-s -w" \
    go build -o todoapp ./cmd/main

FROM alpine:3.20 AS runner
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/todoapp .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

EXPOSE 3000
CMD ["./todoapp"]
