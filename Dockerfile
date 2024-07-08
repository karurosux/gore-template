FROM node:22-alpine AS frontend
WORKDIR /frontend
COPY ./frontend/package.json .
COPY ./frontend/package-lock.json .
RUN npm install
COPY ./frontend .
RUN npm run build
# FE BUILD DONE

FROM golang:1.22.5-alpine AS backend
WORKDIR /backend
COPY ./backend/go.mod .
COPY ./backend/go.sum .
RUN go mod download
COPY ./backend .
RUN go build -o ./dist/gore

FROM alpine:3.14 AS final
WORKDIR /usr/bin/gore
COPY --from=frontend ./frontend/dist/static ./static
COPY --from=backend ./backend/dist/gore .
CMD ["./gore"]
EXPOSE 80



