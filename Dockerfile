FROM golang:1.20 AS builder

ENV GOPROXY='https://goproxy.cn'

WORKDIR /src

COPY . .

# RUN CGO_ENABLED="0" && GOOS=linux && go build -o ./echo ./cmd/echo
RUN go mod download && make build

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /src/echo /app/echo

EXPOSE 8080

ENTRYPOINT ["/app/echo"]