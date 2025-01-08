FROM golang:1.23-alpine3.21 as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY go.* .
RUN go mod download

COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-consignment

# Run container
FROM alpine:3.21.2

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shippy-service-consignment .

CMD ["./shippy-service-consignment"]
