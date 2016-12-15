nginx -g 'daemon off;'

letsencrypt certonly --webroot -w /usr/share/nginx/html -d sutu.shop -d blog.sutu.shop -d mail.sutu.shop -d admin.sutu.shop -d wwwstaging.sutu.shop -d adminstaging.sutu.shop -d api.sutu.shop --email hhnnaahhtt@gmail.com --agree-tos

letsencrypt renew ---quiet --no-self-upgrade
