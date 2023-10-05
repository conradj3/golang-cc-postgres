# Build stage
FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consumer .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/consumer .

CMD ["./consumer"]
