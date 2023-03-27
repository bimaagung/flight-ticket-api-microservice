FROM alpine:latest

RUN mkdir /app

COPY trackApp /app

CMD [ "app/trackApp" ]