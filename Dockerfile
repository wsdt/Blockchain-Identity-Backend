FROM scratch
MAINTAINER Kevin A. Riedl "k.riedl@codeing.io"

# IMPORTANT NOTE:
# Please ensure the Linux binary has been compiled recently and is not deprecated.
# Additionally, you must configure all .json-Files in ./_config/ to make this docker image work.

# Copy binary into image
COPY VID-Card-Backend /

# Copy configs
COPY ./_config /

EXPOSE 8080

CMD ["/VID-Card-Backend"]
