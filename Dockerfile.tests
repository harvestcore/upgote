FROM golang:1.15.5-alpine
LABEL maintainer="Ángel Gómez <agomezm@correo.ugr.es>"
LABEL version="0.2"

WORKDIR /go/src/github.com/harvestcore/HarvestCCode
ENV CGO_ENABLED 0

RUN apk update --no-cache; apk add --no-cache make git

CMD cp -R --symbolic-link /app/test/* /go/src/github.com/harvestcore/HarvestCCode; \
    make test
