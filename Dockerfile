FROM debian:stable-slim

WORKDIR /app
COPY ./conf/logic.yaml /app
COPY ./conf/comet.yaml /app
COPY ./bin /app


EXPOSE 8000
EXPOSE 9000
