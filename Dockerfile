FROM --platform=$BUILDPLATFORM golang:1.22-bookworm AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH}  go build -o todoapp ./cmd/main/main.go

FROM alpine:3.20 AS runner
WORKDIR /app
COPY --from=builder /app/todoapp ./
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates
EXPOSE 3000
ENTRYPOINT ["./todoapp"]
