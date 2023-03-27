FROM alpine:latest

RUN mkdir /app

COPY trackApp /app

EXPOSE 3000

CMD [ "app/trackApp" ]