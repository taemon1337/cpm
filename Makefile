IMAGE=`cat IMAGE`
VERSION=`cat VERSION`
.PHONY: push

build: Dockerfile
      docker build -t $IMAGE .

push: Dockerfile
      docker push $IMAGE
