build:
	go build .

run:
	go build -o ./dist/backend && ./dist/backend

seed:
	go build -o ./dist/backend && ./dist/backend --seed-db

generate-ts:
	go build -o ./dist/backend && ./dist/backend --generate-ts

create-db:
	docker compose up -d

remove-db:
	docker rm -f backend-db

stop-db:
	docker stop backend-db

reset-db:
	docker rm -f backend-db; \
	docker compose up -d; \

reseed:
	docker rm -f backend-db; \
	docker compose up -d; \
	sleep 3; \
	go build -o ./dist/backend && ./dist/backend --seed-db; \