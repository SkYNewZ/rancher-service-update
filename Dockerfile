FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
WORKDIR /
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY rancher-service-update .
ENTRYPOINT ["/rancher-service-update"]
