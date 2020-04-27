FROM golang:1.14 as build
LABEL Maintainer="Soichi Sumi <soichi.sumi@gmail.com>"
ENV GO111MODULE on

COPY . /go/src/tmp
WORKDIR /go/src/tmp
RUN make go-build-amd64

FROM alpine:latest
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.2 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
RUN apk --no-cache add ca-certificates \
    && apk add --no-cache libc6-compat
COPY --from=build /go/src/tmp/exe .
COPY --from=build /go/src/tmp/templates ./templates/
ENTRYPOINT ["./exe"]
