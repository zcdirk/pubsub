FROM golang:alpine AS build
WORKDIR /app
ADD . /app
RUN apk update
RUN apk add build-base
RUN apk add grpc-dev
RUN cd /app && \
    make dependencies && \
    make test && \
    make server

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/bin/pubsub /app

ENV CONFIG=""
ENTRYPOINT ./pubsub --config=$CONFIG
