FROM golang:alpine AS builder

LABEL maintainer="Auditoria de fretes - notifica transportador CTe com erro"

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app
RUN mkdir target

COPY go.mod .
COPY go.sum .
COPY app ./app
COPY domain ./domain
COPY infra ./infra
COPY usecase ./usecase
COPY utils ./utils

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o /app/target/app-go app/main.go 

FROM alpine:3.14.2 AS production

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/target/app-go .
# Expose port 8080 to the outside world
EXPOSE 8080

USER nonroot

# Run the executable
CMD [ "/app/app-go" ]