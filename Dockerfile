FROM golang:alpine AS builder

WORKDIR /app

COPY main.go .

RUN go mod init server

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o server .

FROM scratch

COPY --from=builder /app/server /server

EXPOSE 8080

ENTRYPOINT ["/server"]
