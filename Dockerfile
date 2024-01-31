# builder
FROM ghcr.io/bzhtux/golang:latest AS build-env
LABEL maintainer="Yannick Foeillet <bzhtux@gmail.com>"

ARG GOOS=linux
ARG GOARCH=amd64

WORKDIR /app
RUN mkdir /app/data
RUN mkdir /app/data/db
RUN touch /app/data/db/tsa.db
ADD go.mod go.sum ./
RUN GOOS=${GOOS} GOARCH=${GOARCH} go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -o /tsa ./cmd/main.go

# final image
FROM scratch
LABEL maintainer="Yannick Foeillet <bzhtux@gmail.com>"

WORKDIR /app
COPY --from=build-env /tsa /app/
COPY --from=build-env /app/data /app/data
COPY --from=build-env /app/data/public /app/data/public
USER 1000

EXPOSE 8080

# Run
CMD ["/app/tsa"]