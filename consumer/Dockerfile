# Build stage
FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN GOINSECURE="proxy.golang.org/*,github.com,github.com/*" \
    GONOSUMDB="proxy.golang.org/*,github.com,github.com/*" \
    GOPRIVATE="proxy.golang.org/*,github.com,github.com/*" \
    go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consumer .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/consumer .

EXPOSE 8080

CMD ["./consumer"]