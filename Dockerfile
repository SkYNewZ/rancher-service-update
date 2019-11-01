FROM golang:1.13.4-alpine3.10 as BUILDER
WORKDIR /go/src/github.com/skynewz/rancher-service-update

RUN apk add --update --no-cache \
        git \
        curl \
        ca-certificates

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rancher-service-update .

FROM scratch
WORKDIR /
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=BUILDER /go/src/github.com/skynewz/rancher-service-update/rancher-service-update .
ENTRYPOINT ["/rancher-service-update"]
