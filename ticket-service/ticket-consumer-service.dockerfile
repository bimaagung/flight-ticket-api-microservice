FROM alpine:latest

RUN mkdir /app

COPY ticketConsumerApp /app

EXPOSE 3002

CMD [ "app/ticketConsumerApp" ]