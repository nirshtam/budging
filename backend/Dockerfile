FROM golang:1.26.4-alpine3.23 AS builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./


RUN go build -o app

FROM alpine:3.24.1
WORKDIR /app
COPY --from=builder ./app/app ./
EXPOSE 8000
CMD ["./app"]
