
FROM golang:1.12-stretch AS build-stage
ENV GO111MODULE=on
ADD . /app
RUN cd /app && \
    go build

FROM alpine:3.9
USER root
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-stage /app/octopus-agile-exporter /usr/local/bin/
RUN mkdir /lib64 && \
    ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
    adduser -H -D -s /sbin/nologin -g "Octopus Agile Exporter" exporter exporter
EXPOSE 8080
USER exporter
ENTRYPOINT /usr/local/bin/octopus-agile-exporter
