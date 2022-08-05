# Conspector

Conspector is a project for save and share conspects in real life.
This repo is a backend part of web app. Based on [Gin Framework](https://gin-gonic.com/)

### How to install:

1. install golang 1.18
2. install backend dependencies by command:
```bash
    go mod download
```
3. install redis 7.0 ([instruction](https://redis.io/download/))
4. setup `config.json` in ./configs directory like in example:
```console
{
    "secret_verify_key": "your secret random string ( i use hex sha256 )",
    "database_url": "send me a message so I can send you a test path"
}
```

### How to run:

1. start redis-server:
```bash
    sudo service redis-server start
```
2. start backend(in directory with main.go):
```bash
    go run main.go
```
