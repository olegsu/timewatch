FROM alpine:3.8

RUN apk add --update ca-certificates

COPY dist/linux_386/timewatch /usr/local/bin/

ENTRYPOINT [ "timewatch" ]

CMD [ "--help" ]