FROM alpine:latest

RUN mkdir /app

COPY airplaneApp /app

CMD [ "app/airplaneApp" ]