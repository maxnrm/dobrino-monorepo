.DEFAULT_GOAL=wgo

include .env
export $(shell sed 's/=.*//' .env)

.PHONY: pulldb
pulldb:
	go run scripts/pull_db.go

.PHONY: run
run:
	go run cmd/main.go

.PHONY: build
build:
	go build -o bin/main cmd/main.go

.PHONY: build-run
build-run: build
	./bin/main

.PHONY: build-image
build-image:
	docker build -f build/Dockerfile . -t test-teleflood

.PHONY: run-image
run-image:
	docker run -d --rm -p 8080:8080 --env-file=.env --name test-teleflood test-teleflood:latest

.PHONY: 
deploy-helm:
	helm install tlh -n tlh deploy/helm --set image.tag=${TLH_TAG}

.PHONY: wgo
wgo:
	wgo run cmd/main.go

.PHONY: clean
clean:
	@echo Removing build files
	rm -rf ./bin
	@echo -e "Done.\n"
