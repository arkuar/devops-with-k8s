FROM golang:1.17-alpine as build

WORKDIR /usr/build

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app -ldflags '-w -s'


FROM scratch

WORKDIR /usr/src/app

COPY --from=build /usr/build .

# Add trusted certs
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./app"]