#build stage
FROM golang:1.19-alpine AS builder
ARG GIT_USER
ARG GIT_USER_PASS

RUN apk add --no-cache git build-base
WORKDIR /go/src/app
COPY . .

ENV GOPRIVATE=github.com/pharmatechweb3/*
RUN echo "machine github.com login ${GIT_USER} password ${GIT_USER_PASS}" > ~/.netrc

RUN go build -o /go/bin/main -v ./main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/main /main
ENTRYPOINT /main
EXPOSE 30000
