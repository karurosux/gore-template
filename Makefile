build:
	rm -rf ./dist && \
	mkdir ./dist && \
	cp ./backend/.env ./dist/.env &&\
	cd ./backend && go build -o ../dist/app && cd .. && \
	cd ./frontend && npm run build && cd .. \