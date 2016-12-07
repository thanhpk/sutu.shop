FROM nginx:1.11.5-alpine
MAINTAINER Thanh Pham Kieu

# config nginx
COPY config/nginx.conf /etc/nginx/nginx.conf
COPY config/assetup.conf /etc/nginx/conf.d/assetup.conf

RUN ["rm", "/etc/nginx/conf.d/default.conf"]
RUN ["mkdir", "-p", "/usr/share/nginx/cache/fcgi"]
RUN ["mkdir", "-p", "/srv/assetup/assets"]
