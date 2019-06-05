FROM golang:1.10-alpine3.8 as BUILDER
WORKDIR /go/src/github.com/skynewz/rancher-service-update

RUN apk add --update --no-cache \
        git \
        curl \
        ca-certificates && \
        go get -u github.com/golang/dep/cmd/dep

COPY . .
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rancher-service-update .

FROM scratch
WORKDIR /
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=BUILDER /go/src/github.com/skynewz/rancher-service-update/rancher-service-update .
ENTRYPOINT ["/rancher-service-update"]
