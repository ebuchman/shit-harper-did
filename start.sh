#! /bin/bash

# set git user so we can pull
git config --global user.email "me@me.com"
git config --global user.name "Me"

# if no ssl, turn ssl off
if [ $NO_SSL ]; then
	ln -s /etc/nginx/sites-available/http.conf /etc/nginx/sites-enabled/http.conf
else
	ln -s /etc/nginx/sites-available/https.conf /etc/nginx/sites-enabled/https.conf
fi

/usr/sbin/nginx -g "daemon off;"
