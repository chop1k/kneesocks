FROM golang:1.26 AS build-stage

WORKDIR /src

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o kneesocks cmd/kneesocks/main.go

FROM alpine

RUN mkdir /var/log/kneesocks /etc/kneesocks

COPY --from=build-stage /src/kneesocks /bin/kneesocks

CMD ["kneesocks"]