FROM alpine:latest

RUN mkdir /app

COPY airplaneApp /app

EXPOSE 3001

CMD [ "app/airplaneApp" ]