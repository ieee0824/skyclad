FROM golang:1.10-alpine AS build

RUN apk add --update git

RUN set -e \
    && go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/ieee0824/skyclad

COPY . .

RUN set -e \
    && cd cmd/skyclad \
    && dep ensure \
    && CGO_ENABLED=0 go build .

FROM docker:latest

COPY --from=build /go/src/github.com/ieee0824/skyclad/cmd/skyclad/skyclad /bin/skyclad

ENTRYPOINT ["skyclad"]

CMD ["-lt", "10s", "-it", "10s"]