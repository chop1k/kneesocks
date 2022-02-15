FROM golang as build-stage

WORKDIR /

RUN git clone https://github.com/chop1k/kneesocks

WORKDIR /kneesocks

RUN CGO_ENABLED=0 go test -i -o e2e socks/test/e2e

FROM alpine

COPY --from=build-stage /kneesocks/e2e /bin/e2e

CMD ["e2e"]
