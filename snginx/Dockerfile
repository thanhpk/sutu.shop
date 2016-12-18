from nginx:1.11.6-alpine

maintainer Pham Kieu Thanh

# install certbot
RUN apk update && apk add certbot

COPY certbot.conf /etc/nginx/conf.d/certbot.conf
COPY default.conf default.conf
COPY start.sh start.sh
CMD sh start.sh

