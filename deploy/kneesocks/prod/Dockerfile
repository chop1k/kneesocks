FROM golang:1.26 AS build-stage

WORKDIR /src

COPY go.mod go.sum .

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 \
    go build -trimpath -ldflags="-s -w" -o kneesocks cmd/kneesocks/main.go

FROM alpine:3.24.1

RUN mkdir /var/log/kneesocks /etc/kneesocks

COPY --from=build-stage /src/kneesocks /bin/kneesocks

ENTRYPOINT ["/bin/kneesocks"]