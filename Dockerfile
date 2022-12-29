FROM golang:1.18-alpine as builder

WORKDIR /srv

COPY . /srv

RUN go build -o main .

FROM alpine:3.15

COPY --from=builder /srv/main /srv/main
COPY --from=builder /srv/certs /certs

ENTRYPOINT ["/srv/main"]