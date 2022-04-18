FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./

COPY *.go ./

RUN go build -o /large-file-processing

EXPOSE 8080
EXPOSE 8081

CMD [ "/large-file-processing" ]