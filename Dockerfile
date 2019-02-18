FROM scratch
MAINTAINER Kevin A. Riedl "k.riedl@codeing.io"

# Copy binary into image
COPY VID-Card-Backend /

# Copy configs
COPY ./_config /

CMD ["/VID-Card-Backend"]
