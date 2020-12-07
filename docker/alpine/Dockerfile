FROM alpine:3.12

ARG DRONE_ARCHIVE_VERSION=0.0.0
ENV DRONE_ARCHIVE_VERSION=${DRONE_ARCHIVE_VERSION}

LABEL version="${DRONE_ARCHIVE_VERSION}" \
    description="Drone plugin that provides Tar/GZ and Zip archiving with optional checksum computation." \
    maintainer="Jose Quintana <joseluisq.net>"

COPY ./docker/alpine/entrypoint.sh /
COPY release/linux/amd64/drone-archive /usr/local/bin/

ENTRYPOINT ["/entrypoint.sh"]
CMD ["drone-archive"]

# Metadata
LABEL org.opencontainers.image.vendor="Jose Quintana" \
    org.opencontainers.image.url="https://github.com/joseluisq/drone-archive" \
    org.opencontainers.image.title="Drone Archive Plugin" \
    org.opencontainers.image.description="Drone plugin that provides Tar/GZ and Zip archiving with optional checksum computation." \
    org.opencontainers.image.version="${DRONE_ARCHIVE_VERSION}" \
    org.opencontainers.image.documentation="https://github.com/joseluisq/drone-archive"
