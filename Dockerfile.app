# builder
FROM golang:alpine AS build-env
LABEL maintainer="Yannick Foeillet <yfoeillet@vmware.com>"

ARG GOOS=linux
ARG GOARCH=amd64

# wokeignore:rule=he/him/his
RUN apk --no-cache add build-base git mercurial gcc curl
RUN mkdir -p /go/src/github.com/bzhtux/tsa
ADD . /go/src/github.com/bzhtux/tsa/
WORKDIR /go/src/github.com/bzhtux/tsa
RUN GOOS=${GOOS} GOARCH=${GOARCH} go get ./...
RUN GOOS=${GOOS} GOARCH=${GOARCH} go build -o tsa cmd/main.go


# final image
FROM alpine
# FROM scratch
LABEL maintainer="Yannick Foeillet <yfoeillet@vmware.com>"

# wokeignore:rule=he/him/his
RUN apk --no-cache add curl jq tini
RUN adduser -s /bin/sh -u 10000 -D app
RUN mkdir -p /goss
COPY goss.yaml /goss
WORKDIR /app
RUN mkdir /app/data
RUN mkdir /app/data/db
COPY data/public /app/data/public
RUN chown -R app /app
COPY --from=build-env /go/src/github.com/bzhtux/tsa/tsa /app/
USER 0
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/tsa"]