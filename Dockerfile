FROM golang:1.22-alpine AS builder

RUN apk add --no-cache bash

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o todoapp ./cmd/main

FROM alpine:3.20 AS runner

WORKDIR /app

COPY --from=builder /app/todoapp .

COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

EXPOSE 3000

CMD ["./todoapp"]
