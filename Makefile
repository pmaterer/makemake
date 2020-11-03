test:
	go test -v ./...

build:
	docker-compose build makemake

run:
	docker-compose up -d makemake

restart:
	docker-compose kill makemake
	docker rm makemake
	docker-compose up -d makemake

stop:
	docker stop makemake
	docker rm makemake

db:
	docker-compose up -d db pgadmin