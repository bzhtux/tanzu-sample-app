# builder
FROM ghcr.io/bzhtux/golang:latest AS build-env
# FROM golang:alpine AS build-env
LABEL maintainer="Yannick Foeillet <bzhtux@gmail.com>"

# ARG GOOS=linux
# ARG GOARCH=amd64

ARG GOOS=darwin
ARG GOARCH=arm64

WORKDIR /app
RUN apk --no-cache add build-base git gcc sqlite-dev clang
RUN mkdir /app/data
RUN mkdir /app/data/db
RUN touch /app/data/db/tsa.db
ADD go.mod go.sum ./
RUN go env -w CGO_ENABLED=1
# RUN GOOS=${GOOS} GOARCH=${GOARCH} go mod download
RUN go mod download

COPY . ./

# RUN GOOS=${GOOS} GOARCH=${GOARCH} go build -o /tsa ./cmd/main.go
RUN go build -o /app/tsa ./cmd/main.go
RUN echo $files


# final image
FROM scratch
# FROM alpine
LABEL maintainer="Yannick Foeillet <bzhtux@gmail.com>"

WORKDIR /app
COPY --from=build-env /app/tsa /app/
COPY --from=build-env /app/data /app/data
# COPY --from=build-env /app/data/public /app/data/public
USER 1000

EXPOSE 8080

# Run
CMD ["/app/tsa"]