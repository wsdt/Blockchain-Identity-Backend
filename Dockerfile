FROM alpine:3.6 as alpine

# IMPORTANT NOTE:
# Additionally, you must configure all .json-Files in ./_config/ to make this docker image work.

MAINTAINER Kevin A. Riedl "k.riedl@codeing.io"

RUN apk add -U --no-cache ca-certificates

FROM scratch
ENTRYPOINT []
WORKDIR /
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


# Copy binary into image
COPY VID-Card-Backend /

# Copy configs
COPY ./_config /_config


CMD ["/VID-Card-Backend"]
