FROM golang:1.16-alpine3.12 as builder
WORKDIR /src
COPY . .
RUN go build -o server

FROM alpine:3.12
WORKDIR /app
COPY --from=builder /src/server .
ENTRYPOINT ["/app/server"]
