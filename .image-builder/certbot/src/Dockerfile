ARG GOLANG_IMAGE_VERSION="latest"
ARG BASE_IMAGE_VERSION="latest"
FROM golang:$GOLANG_IMAGE_VERSION AS builder

# Copy the entrypoint Go script into the builder
# COPY src/entrypoint.go .
# COPY src/go.mod .

# Build the Go binary
# RUN go build -o /usr/local/bin/entrypoint entrypoint.go

# Stage 2: Create the final image
FROM certbot/dns-cloudflare:$BASE_IMAGE_VERSION

# Create necessary directories
RUN mkdir -p /etc/letsencrypt /var/lib/letsencrypt /cloudflared

# Copy the compiled Go binary from the builder stage
# COPY --from=builder /usr/local/bin/entrypoint /usr/local/bin/entrypoint

# Define volumes
VOLUME ["/etc/letsencrypt", "/var/lib/letsencrypt", "/cloudflared"]

# Set environment variables for domain and email (can also be overridden at runtime)
ARG DOMAIN_NAME
ARG EMAIL_ADDRESS
ENV DOMAIN_NAME=$DOMAIN_NAME
ENV EMAIL_ADDRESS=$EMAIL_ADDRESS

COPY bin/entrypoint /usr/local/bin/entrypoint

# Set the ENTRYPOINT to the compiled Go binary
ENTRYPOINT ["/usr/local/bin/entrypoint"]
