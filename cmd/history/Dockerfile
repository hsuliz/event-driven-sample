FROM golang:1.24.1-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o history ./cmd/history

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app .
ENTRYPOINT ["./history"]
