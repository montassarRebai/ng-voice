FROM debian
COPY ./go-app /app
ENTRYPOINT /app
ENTRYPOINT ["tail", "-f", "/dev/null"]
