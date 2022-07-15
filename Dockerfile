FROM golang:1.15.5-alpine as build

WORKDIR /go/src/github.com/harvestcore/upgote
ENV CGO_ENABLED 0
COPY . .
RUN apk update --no-cache; apk add --no-cache make git; make buildapp


FROM alpine:3.7

COPY --from=build /go/src/github.com/harvestcore/upgote/upgote .
COPY --from=build /go/src/github.com/harvestcore/upgote/Makefile .
EXPOSE 80
RUN apk update --no-cache; apk add --no-cache make

CMD make start

HEALTHCHECK CMD curl --fail http://localhost/api/healthcheck || exit 1