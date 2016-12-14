from nginx:1.11.6-alpine

maintainer Pham Kieu Thanh

# install certbot
RUN apk update && apk add ca-certificates wget && update-ca-certificates
RUN wget https://dl.eff.org/certbot-auto
RUN chmod a+x certbot-auto

RUN ./certbot-auto

RUN mkdir -p /srv/sutu.shop
RUN ./certbot-auto certonly --webroot -/srv/sutu.shop -d sutu.shop -d blog.sutu.shop -d mail.sutu.shop -d admin.sutu.shop -d www.staging.sutu.shop -d admin.staging.sutu.shop -d api.sutu.shop
RUN ./certbot-auto renew --dry-run
RUN ./certbot-auto renew --quite --no-self-upgrade
