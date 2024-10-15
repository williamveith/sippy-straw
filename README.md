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
sippystraw
├── README.md
├── appserver
│   └── docker-compose.yml
└── straw
    ├── Buildfiles
    │   ├── bin
    │   │   └── main
    │   ├── go.mod
    │   ├── go.sum
    │   └── main.go
    ├── Dockerfiles
    │   ├── Dockerfile.certbot
    │   └── Dockerfile.cloudflare
    ├── Makefile
    ├── certbot
    │   ├── entrypoint.go
    │   └── go.mod
    ├── cloudflared
    ├── docker-compose.yml
    └── nginx
        └── app.conf
```
