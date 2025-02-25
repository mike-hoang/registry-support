# Index Server build stage
FROM golang:1.18-alpine AS index-builder
WORKDIR /tools
COPY . .
RUN CGO_ENABLED=0 go build -mod=vendor -o index-server main.go

# Application image
FROM registry.access.redhat.com/ubi8/ubi-minimal AS runner
USER root

# Install and configure dependencies
RUN microdnf update -y && microdnf install shadow-utils findutils && rm -rf /var/cache/yum
COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh

# Copy index server
COPY --from=index-builder /tools/index-server /registry/index-server
RUN chgrp -R 0 /registry && \
    chmod -R g=u /registry

# Create a non-root user to run the server as
RUN set -x ; \
    adduser www-data -u 1001 -G root && exit 0

# Modify the permissions on the necessary files to allow the container to properly run as a non-root UID
RUN mkdir -p /www/data && chmod -R g+rwx /www/data

# Set env vars for the locations of the devfile stacks and index.json
ENV DEVFILE_STACKS /registry/stacks
ENV DEVFILE_SAMPLES /registry/samples
ENV DEVFILE_INDEX /registry/index.json
ENV DEVFILE_BASE64_INDEX /www/data/index_base64.json
ENV DEVFILE_SAMPLE_INDEX /www/data/sample_index.json
ENV DEVFILE_SAMPLE_BASE64_INDEX /www/data/sample_base64_index.json
ENV DEVFILE_STACK_INDEX /www/data/stack_index.json
ENV DEVFILE_STACK_BASE64_INDEX /www/data/stack_base64_index.json

USER 1001

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
