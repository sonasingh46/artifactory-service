#
# This Dockerfile builds a recent artifactory-service using the latest binary from
# artifactory-service  releases.
#
FROM alpine:3.11.5

RUN apk add --no-cache \
    iproute2 \
    bash \
    curl \
    net-tools \
    mii-tool \
    procps \
    libc6-compat \
    ca-certificates

COPY artifactory-service /usr/local/bin/artifactory-service

ARG ARCH
ARG DBUILD_DATE
ARG DBUILD_REPO_URL

LABEL org.label-schema.name="artifactory-service"
LABEL org.label-schema.description="Jfrog artifactory proxy service"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.build-date=$DBUILD_DATE
LABEL org.label-schema.vcs-url=$DBUILD_REPO_URL

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/artifactory-service"]

