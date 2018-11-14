version=$(shell cat version)

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o postgres-go-makedb

docker: build
	docker build . -t craftypenguins/postgres-go-makedb:$(version)

push: docker
	docker push craftypenguins/postgres-go-makedb:$(version)
