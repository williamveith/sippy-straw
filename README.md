# Pipe

## Adding Credentials

Replace the files listed in cloudflared/manifest.md
Build specialized images with .image-builder/Makefile
Run docker-compose to start

```sh
docker-compose up --build -d
```

## Project File List

```txt
pipe
├── .env
├── .gitignore
├── .image-builder
│   ├── .env
│   ├── Makefile
│   ├── README.md
│   ├── bin
│   │   └── main
│   ├── cosign.key
│   ├── cosign.pub
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── src
│       └── Dockerfile.cloudflare
├── .vscode
│   ├── launch.json
│   └── settings.json
├── README.md
├── certbot
│   └── Dockerfile
├── cloudflared
│   ├── 60f2a096-47a8-405a-af97-f5088a41231d.json
│   ├── cert.pem
│   ├── cloudflare.ini
│   ├── config.yml
│   └── manifest.md
├── docker-compose.yml
└── nginx
    └── app.conf
```
