FROM golang:1.19-alpine AS builder

ENV GOPATH /go
ENV GO111MODULE on
ENV CGO_ENABLED 0

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:3.12
WORKDIR /app
COPY --from=builder /app /app

EXPOSE 8080

CMD ["./main"]
