FROM alpine:latest

RUN mkdir /app

COPY mailServiceApp /app
COPY template /templates
CMD [ "/app/mailServiceApp"]