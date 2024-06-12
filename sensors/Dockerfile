#Builder image
FROM golang:alpine3.15 as builder
RUN apk add git
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o rainbow-hat-sensor-service .

#Runner image
FROM alpine:3.15
COPY --from=builder /build/rainbow-hat-sensor-service .

ENTRYPOINT [ "./rainbow-hat-sensor-service" ]