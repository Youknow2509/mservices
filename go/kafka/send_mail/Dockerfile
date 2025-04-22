# Build stage
FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.mservice.sendmail.com ./main.go

# Final stage 
# Use alpine as base image for runtime
FROM alpine:latest  

# Install CA certificates and necessary dependencies
RUN apk add --no-cache ca-certificates

COPY ./config /config
COPY ./html-template /html-template

COPY --from=builder /build/crm.mservice.sendmail.com /

ENTRYPOINT [ "/crm.mservice.sendmail.com", "config/local.yaml" ]
