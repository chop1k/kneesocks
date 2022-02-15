FROM golang as build-stage



FROM archlinux

COPY --from=build-stage /kneesocks /bin/kneesocks

CMD ["kneesocks"]