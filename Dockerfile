FROM ubuntu:latest
MAINTAINER Coin Culture <support@coinculture.info>

RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-get dist-upgrade -y && apt-get install -y git nginx  && apt-get clean all

ENV repo /srv/www
RUN mkdir -p $repo
COPY . $repo/
RUN chown -R www-data:www-data $repo
WORKDIR $repo

RUN rm /etc/nginx/sites-enabled/*
RUN rm /etc/nginx/sites-available/*
COPY nginx/* /etc/nginx/sites-available/

EXPOSE 80
EXPOSE 443

COPY ./start.sh /start.sh
RUN chmod 755 /start.sh

CMD ["/start.sh"]
