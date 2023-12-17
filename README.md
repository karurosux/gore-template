# Type Safe Frontend+Backend Template - React + Go + PostgreSql Template

Base application built with the next stack:

- React: Using redux toolkit and redux toolkit query to fetch data from the backend, the application was generated with vite.
- Go: Using golang in the backend, echo for the API framework and some other dependecies for the creating utilities.
- PostgreSql: The DB is created in docker, see the makefile inside the backend folder and docker-compose.yml file in the same folder for more detail.

## Prerequesites

1. Unix OS(MacOs, Linux or WSL).
2. Docker and Docker Compose if applicable.
3. NodeJs(Suggest to use NVM).
4. Make command.

## Getting Started

1. Run backend:
   1. Move to backend folder in terminal (`cd ./backend`).
   2. Run command `make create-db` for creating db container in docker.
   3. Run command `make seed` for loading development data in db.
   4. Run command `make run` to run backend application.
2. Run frontend:
   1. Move to frontend folder in terminal(`cd ./frontend`).
   2. Run command `npm install` to install dependencies.
   3. Run command `npm start` to run application in development mode.
3. Open `http://locahost:8080` in a web browser.

## Generating Frontend Types

This template generates types for the frontend from the backend types, the configuration for this can be found at
`./backend/ts-generator.go` file.

These types get generated into `./frontend/src/model/generated` folder.

For regenerating these types just move to
`./backend` in terminal and run the next command:

```
make generate-ts
```

## Building Application Production

Of course this may need more work to be done here, but this template gets a basic approach for this process ready, for generating a prod version just run the next command from the root folder:

```
make build
```

This will generate a folder `./dist` containing the backend binary, the frontend static folder containg static files from our application and a dot env file.

The static files will be served by our backend in the root path, so frontend in this case it would be served in the same port.

To execute the binary just run the next command:

```
cd ./dist && ./app
```
