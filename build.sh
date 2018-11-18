#!/bin/bash

GOVERSION=1.10
PROJECT=github.com/taemon1337/cpm

if [ -z "$GOROOT" ]; then
  GOROOT=$HOME/code/go
fi

cmd="go build -ldflags -w -a $PROJECT"

if [ ! -z "$1" ]; then
  cmd="${@}"
fi

docker run --rm -it \
  --user $UID:$UID \
  -e GOOS=linux \
  -e GOARCH=amd64 \
  -e CGO_ENABLED=0 \
  -v $GOROOT:/go \
  -v $HOME/.cache:/.cache \
  -w /go/src/$PROJECT \
  golang:$GOVERSION \
  $cmd

