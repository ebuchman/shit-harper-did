#! /bin/bash
set -e

# set git user so we can pull
git config user.email "me@me.com"
git config user.name "Me"

# run the shd-server
./shd-server
