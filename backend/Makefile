build:
	go build .

run:
	go build -o ./dist/app && ./dist/app

seed:
	go build -o ./dist/app && ./dist/app --seed-db

generate-ts:
	go build -o ./dist/app && ./dist/app --generate-ts

create-db:
	docker compose up -d

remove-db:
	docker rm -f app-db

stop-db:
	docker stop app-db

reset-db:
	docker rm -f app-db; \
	docker compose up -d; \

reseed:
	docker rm -f app-db; \
	docker compose up -d; \
	sleep 3; \
	go build -o ./dist/app && ./dist/app --seed-db; \