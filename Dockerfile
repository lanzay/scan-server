# Stage 1 - the build www
FROM node:14-alpine as build-deps
WORKDIR /usr/src/app

COPY ./servers/www/site_v1/package.json ./
COPY ./servers/www/site_v1/src /usr/src/app/src
COPY ./servers/www/site_v1/public /usr/src/app/public

RUN npm install
#RUN npm audit fix
RUN npm run build


# Stage 2 - the production environment
FROM golang:1.17
WORKDIR /go/src/app
COPY . .
COPY --from=build-deps /usr/src/app/build ./servers/www/site_v1/build

RUN go get -d -v ./...
RUN go build -v ./cmd/scan-server

USER 1001
EXPOSE 3030

CMD ["./scan-server"]
