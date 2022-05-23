GOPATH ?= $(HOME)/go

up:
	docker-compose up --force-recreate --build

down:
	docker-compose down