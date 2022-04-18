# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go get github.com/TranManhChung/large-file-processing/parser
RUN go get github.com/TranManhChung/large-file-processing/pkg/util
RUN go get github.com/TranManhChung/large-file-processing/quering
RUN go get github.com/TranManhChung/large-file-processing/storage

COPY *.go ./

RUN go build -o /large-file-processing

EXPOSE 8080

CMD [ "/large-file-processing" ]