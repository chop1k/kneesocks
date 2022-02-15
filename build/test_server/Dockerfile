FROM golang as build-stage

WORKDIR /

RUN git clone https://github.com/chop1k/kneesocks

WORKDIR /kneesocks

RUN CGO_ENABLED=0 go build -o test_server cmd/test_server/main.go

FROM alpine

COPY --from=build-stage /kneesocks/test_server /bin/test_server

RUN mkdir /etc/kneesocks /usr/share/kneesocks /usr/share/kneesocks/images

CMD ["test_server"]
