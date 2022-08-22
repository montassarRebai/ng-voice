FROM debian
COPY ./go-app /app
ENTRYPOINT /app
