FROM golang:latest

WORKDIR /app
COPY ./app /app

make

CMD ["make run"]
