FROM --platform=$BUILDPLATFORM golang:1.22-bookworm AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY . .
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH}  go build -o bin/todoapp ./cmd/main/main.go

FROM alpine:3.20 AS runner
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/bin/todoapp .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates
EXPOSE 3000
CMD ["./bin/todoapp"]
