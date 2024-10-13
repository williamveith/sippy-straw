FROM certbot/dns-cloudflare:v2.11.0@sha256:43bf5ab075f5018ae4a4a1ec9021151dab0294d72898f887c38ec86d34e4f7d3

# Create necessary directories
RUN mkdir -p /etc/letsencrypt

# Define volumes
VOLUME ["letsencrypt:/etc/letsencrypt"]

# Set the entrypoint to check if the certificate exists, request a new one if not, otherwise renew
ENTRYPOINT ["/bin/sh", "-c", "trap exit TERM; while :; do if [ ! -f /etc/letsencrypt/live/dropbox.williamveith.com/fullchain.pem ]; then echo 'Certificate not found, requesting a new one...'; certbot certonly --non-interactive --agree-tos --dns-cloudflare --dns-cloudflare-credentials /etc/cloudflare/cloudflare.ini --dns-cloudflare-propagation-seconds 120 -d dropbox.williamveith.com --email w.il.lia.mveith+kxnfxl6l@gmail.com; else echo 'Certificate found, attempting renewal...'; certbot renew --cert-name dropbox.williamveith.com --dns-cloudflare --dns-cloudflare-credentials /etc/cloudflare/cloudflare.ini --dns-cloudflare-propagation-seconds 120; fi; sleep 12h & wait $!; done;"]
