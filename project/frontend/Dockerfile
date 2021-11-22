FROM node:16-alpine as build

WORKDIR /usr/src/app

COPY . .

RUN npm install && \
    npm run build

FROM node:16-alpine

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/build .

RUN npm install -g serve && \
    adduser -D appuser

USER appuser

CMD ["serve", "-s"]