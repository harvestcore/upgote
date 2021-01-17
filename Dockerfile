FROM golang:alpine
LABEL maintainer="Ángel Gómez <agomezm@correo.ugr.es>" version="0.3"

WORKDIR /app/test
ENV CGO_ENABLED 0

RUN adduser -D hcc && addgroup -S hcc hcc

RUN apk update --no-cache; apk add --no-cache make git

ENV GOPATH=/home/hcc/go

USER hcc

CMD make test
