FROM golang:1.26 AS build-stage

WORKDIR /src

COPY go.mod go.sum .

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 \
    go build -trimpath -ldflags="-s -w" -o test_server cmd/test_server/main.go

FROM alpine:3.24.1

RUN mkdir /etc/kneesocks /usr/share/kneesocks

COPY --from=build-stage /src/test_server /bin/test_server

USER nobody:nogroup

ENTRYPOINT ["/bin/test_server"]
