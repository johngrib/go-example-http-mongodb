FROM golang:1.12

WORKDIR $HOME/go-example-http-mongodb
COPY . .

RUN go mod tidy
RUN go build

EXPOSE 8080
CMD ["go-example-http-mongodb"]
