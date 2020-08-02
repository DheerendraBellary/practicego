FROM alpine:3.5

ENV WEBSERVER_PORT 8080
RUN mkdir -p /app
COPY ./simplehttp /app/simplehttp
RUN chmod +x /app/simplehttp

EXPOSE ${WEBSERVER_PORT}
ENTRYPOINT /app/simplehttp


