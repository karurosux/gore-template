build:
	sudo docker build -t gore-template .

images:
	sudo docker images

containers:
	sudo docker ps -a

run:
	sudo docker run --name gore -d gore-template

gore-logs:
	sudo docker container logs gore

add-containers:
	sudo docker-compose up -d

remove-containers:
	sudo docker rm gore --force
	sudo docker rm gore-db --force

drop:
	make remove-containers
	sudo docker network rm database --force

reset:
	make drop
	make install

install:
	sudo docker network create database
	make build
	make add-containers

seed:
	cd ./backend && go build -o ./dist/gore && ./dist/gore --seed-db && cd ..
