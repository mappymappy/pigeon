pigeon [![GoDoc](http://godoc.org/github.com/mappymappy/pigeon?status.svg)](http://godoc.org/github.com/mappymappy/pigeon)

## Introduciton

ChatServer using gRPC & golang.

## Install

```
 go get github.com/mappymappy/pigeon
```

## Configuration & Usage

pigeon works depending on environment variables.
see `.env.sample`.
Here is an example of when using [envrc](https://github.com/direnv/direnv).

```
 cp .env.sample .envrc
 direnv allow
```

* Server

```
make run-sever
```

* Client

```
make run-client1
make run-client2
```

## Author
[marnie_ms4](https://github.com/mappymappy?tab=repositories)
