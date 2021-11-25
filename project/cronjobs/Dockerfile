FROM alpine:3.15

WORKDIR /usr/src/job

COPY daily.sh .

RUN apk --no-cache add curl && \
    chmod +x daily.sh

CMD [ "./daily.sh" ]