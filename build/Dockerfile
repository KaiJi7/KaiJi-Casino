FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app .

FROM alpine

WORKDIR /app
COPY --from=builder /app/app ./
COPY --from=builder /etc/ssl /etc/ssl

COPY ./configs/config.yaml ./configs/

# Command to run the executable
CMD ["./app"]
