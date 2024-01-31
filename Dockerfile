# builder
FROM ghcr.io/bzhtux/golang:latest AS build-env
LABEL maintainer="Yannick Foeillet <bzhtux@gmail.com>"

WORKDIR /app

ADD go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tsa

# final image
FROM scratch
LABEL maintainer="Yannick Foeillet <bzhtux@gmail.com>"

WORKDIR /app
COPY --from=build-env /tsa /app/
RUN mkdir /app/data
RUN mkdir /app/data/db
COPY data/public /app/data/public
USER 1000

EXPOSE 8080

# Run
CMD ["/app/tsa"]