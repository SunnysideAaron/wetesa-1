FROM golang:1.23.6-bookworm
# FROM bitnami/golang:1.23-debian-12

WORKDIR /

# RUN go clean -modcache

COPY . .

WORKDIR /common

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.29.0
RUN go install github.com/fdietze/sqlc-gen-from-template@v1.0.0