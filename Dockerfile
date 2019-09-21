

FROM golang:latest

COPY ./project /go/src/

COPY entrypoint.sh entrypoint.sh

ENTRYPOINT ["sh", "entrypoint.sh"]

