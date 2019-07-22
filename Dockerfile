FROM golang:1.11-alpine AS builder

RUN apk add bash ca-certificates git

WORKDIR /application

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

# Copy all files in currend directiry into home directory
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./bin/app .

FROM alpine:3.9
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /application
COPY --from=builder /application /application

ENV GO111MODULE=on \
    TAX_JAR_TOKEN="" \
    ZIP_CODE_FILE="" \
    CACHE_PATH="./cache" \
    MAX_RPS=250

ENTRYPOINT /application/bin/app