FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN cd zipcode-gateway/cmd/zipcode-gateway && go build -o /app/zipcode-gateway

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/zipcode-gateway .
EXPOSE 8080
CMD ["./zipcode-gateway"]
