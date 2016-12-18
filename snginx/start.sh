
rm /etc/nginx/conf.d/default.conf
nginx -g 'daemon off;' 1>>output.log 2>>error.log

letsencrypt certonly --webroot -w /usr/share/nginx/html -d sutu.shop -d blog.sutu.shop -d mail.sutu.shop -d admin.sutu.shop -d wwwstaging.sutu.shop -d adminstaging.sutu.shop -d api.sutu.shop --email hhnnaahhtt@gmail.com --agree-tos 1>>output.log 2>>error.log

letsencrypt renew --quiet --no-self-upgrade 1>>output.log 2>>error.log
