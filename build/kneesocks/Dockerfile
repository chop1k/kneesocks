FROM golang as build-stage

WORKDIR /

RUN git clone https://github.com/chop1k/kneesocks

WORKDIR /kneesocks

RUN CGO_ENABLED=0 go build -o kneesocks cmd/kneesocks/main.go

FROM alpine

COPY --from=build-stage /kneesocks/kneesocks /bin/kneesocks

RUN mkdir /var/log/kneesocks /etc/kneesocks

CMD ["kneesocks"]