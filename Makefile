BINARY=argocd-tf-plugin

default: build

build:
	go build -o ${BINARY} .

build-docker:
	docker build --target=prod . 

install: build