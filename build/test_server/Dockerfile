FROM golang:1.26 AS build-stage

WORKDIR /src

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o test_server cmd/test_server/main.go

FROM alpine

RUN mkdir /etc/kneesocks /usr/share/kneesocks

COPY --from=build-stage /src/test_server /bin/test_server

CMD ["test_server"]
