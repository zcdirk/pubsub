FROM golang:alpine AS build
WORKDIR /app
ADD . /app
RUN apk update && apk add make protoc
RUN cd /app && \
    make dependencies && \
    make server

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/pubsub-server /app

EXPOSE 7476
ENTRYPOINT ./pubsub-server
