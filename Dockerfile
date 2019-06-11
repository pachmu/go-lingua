FROM golang:1.12

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y libgtk-3-dev

COPY . /src/go-lingua
