FROM golang:1.17

RUN mkdir -p /home/app

COPY . /home/app

WORKDIR /home/app
RUN go mod download
RUN go build -o /hello

CMD ["/hello"]



