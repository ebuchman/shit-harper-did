FROM golang:1.4
MAINTAINER Coin Culture <support@coinculture.info>

ENV user shd

RUN groupadd -r $user \
  && useradd -r -s /bin/false -g $user $user

USER $user

ENV repo /go/src/github.com/ebuchman/shit-harper-did
RUN mkdir -p $repo
COPY . $repo/
WORKDIR $repo
RUN go build -o ./shd-server ./server

EXPOSE 8080

COPY ./start.sh /start.sh

CMD ["/start.sh"]
