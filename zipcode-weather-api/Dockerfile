FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN cd zipcode-weather-api/cmd/zipcode-weather-api && go build -o /app/zipcode-weather-api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/zipcode-weather-api .
EXPOSE 8081
CMD ["./zipcode-weather-api"]
