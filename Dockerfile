# BUILD
FROM golang:1.19.4-alpine3.17 AS build-env

ENV APP_NAME spotiyou
ENV CMD_PATH cmd/spotiyou/spotiyou.go

WORKDIR /$APP_NAME
COPY ./ ./
RUN go build $CMD_PATH

# RUN
FROM alpine:3.17.0 AS prod-env

ENV APP_NAME spotiyou

COPY --from=build-env /$APP_NAME .
CMD ./$APP_NAME