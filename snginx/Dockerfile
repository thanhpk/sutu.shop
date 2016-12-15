from nginx:1.11.6-alpine

maintainer Pham Kieu Thanh

# install certbot
RUN apk update && apk add ca-certificates certbot && update-ca-certificates

COPY default.conf /etc/nginx/conf.d/default.conf
COPY start.sh start.sh
CMD sh start.sh

