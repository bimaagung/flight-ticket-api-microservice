FROM alpine:latest

RUN mkdir /app

COPY ticketApp /app

CMD [ "app/ticketApp" ]