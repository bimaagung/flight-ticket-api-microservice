FROM alpine:latest

RUN mkdir /app

COPY ticketApp /app

EXPOSE 3002

CMD [ "app/ticketApp" ]