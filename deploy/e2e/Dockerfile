FROM golang:1.26 AS build-stage

WORKDIR /src

COPY go.mod go.sum .

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go test -c -o e2e socks/test/e2e

FROM alpine:3.24.1

COPY --from=build-stage /src/e2e /bin/e2e

ENTRYPOINT ["/bin/e2e"]
