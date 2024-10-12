# Sensitive Files Manifest

The following files are required for the Cloudflare tunnel setup and must be placed in the this directory. These files are ignored in version control for security reasons.

- **`60f2a096-47a8-405a-af97-f5088a41231d.json`**: Tunnel credential, required to authenticate the tunnel to Cloudflare.
- **`cert.pem`**: Cloudflare origin server certificate, used for secure connections.
- **`config.yml`**: Cloudflared configuration, specifies how Cloudflared should be set up and what services to expose.
- **`cloudflare.ini`**: Cloudflare API credentials for managing DNS or other services.

Make sure these files are available before running the project.
