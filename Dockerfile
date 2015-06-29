FROM golang:1.4
MAINTAINER Coin Culture <support@coinculture.info>

ENV USER shd

RUN groupadd -r $USER \
  && useradd -r -s /bin/false -g $USER $USER

ENV repo /go/src/github.com/ebuchman/shit-harper-did
RUN mkdir -p $repo
COPY . $repo/
WORKDIR $repo
RUN go build -o ./shd-server ./server

RUN chown -R $USER:$USER .
USER $USER

EXPOSE 8080

CMD ["./start.sh"]
