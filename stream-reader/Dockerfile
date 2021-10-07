FROM golang:1.12 AS build
COPY . /app
WORKDIR /app
RUN go get -d
RUN go build -o main .