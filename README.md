# go-chat
## Description
goによるチャットアプリ用のapi server

## Build environment
### How to start at local environment
```

# Install binary of go module
$ docker-compose run --rm app make install　

$ docker-compose up -d
```
### Build and run of production version
```
docker build -t go-chat .
docker run --name go-chat -d -it -p 8080:8080  go-chat
```

