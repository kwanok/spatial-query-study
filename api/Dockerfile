FROM golang:1.19-alpine AS builder

WORKDIR /app

ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY . .

RUN go mod tidy

RUN go build -o main main.go

FROM alpine:latest

WORKDIR /app

RUN apk --update add ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/runtime-config-live.yaml .

EXPOSE 8080

CMD ["/app/main", "-config=/app/runtime-config-live.yaml"]