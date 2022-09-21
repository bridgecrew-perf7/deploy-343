FROM nginx:1.20-alpine

RUN mv /etc/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf.bk
COPY fileserver.cnf /etc/nginx/conf.d/fileserver.conf

EXPOSE 80
