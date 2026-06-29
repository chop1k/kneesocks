FROM golang:1.26 AS build-stage

WORKDIR /src

# .dockerignore
# compilation flags
# buildkit cache mounts

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -c -o e2e socks/test/e2e

FROM alpine

COPY --from=build-stage /src/e2e /bin/e2e

CMD ["e2e"]
