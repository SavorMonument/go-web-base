FROM golang:1.23-bookworm

WORKDIR /app

COPY ./ui ./ui
COPY ./config ./config

COPY ./target/web /app/web
# Fix permissions for executable
RUN chmod +x ./web

EXPOSE 80

CMD ["./web"]
