FROM golang:1.20

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN make build

EXPOSE 6500

CMD ["./ms-api"]
