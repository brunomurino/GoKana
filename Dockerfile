

FROM golang:latest

COPY ./project /go/src/

ENTRYPOINT ["sh", "/go/src/entrypoint.sh"]

