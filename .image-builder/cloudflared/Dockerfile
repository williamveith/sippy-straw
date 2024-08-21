# Start from an Alpine Linux base image
FROM alpine:3.20 AS builder

ARG CLOUDFLARED_VERSION

# Install necessary tools and download the cloudflared binary
RUN apk add --update --no-cache wget binutils && \
    wget -q https://github.com/cloudflare/cloudflared/releases/download/${CLOUDFLARED_VERSION}/cloudflared-linux-arm64 -O /usr/local/bin/cloudflared && \
    strip --strip-unneeded /usr/local/bin/cloudflared

# Change ownership of the cloudflared binary to the non-root user
RUN addgroup -S cloudflare && adduser -S -G cloudflare cloudflare && \
    chown cloudflare:cloudflare /usr/local/bin/cloudflared && \
    chmod +x /usr/local/bin/cloudflared && \
    apk del wget binutils

# Minimal setup with scratch as base
FROM scratch
COPY --from=builder /usr/local/bin/cloudflared /usr/local/bin/cloudflared
COPY --from=builder /etc/passwd /etc/group /etc/
USER cloudflare:cloudflare
ENTRYPOINT ["/usr/local/bin/cloudflared"]