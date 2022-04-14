# docker multistage build
FROM nginx:1.20-alpine

# npm run-script build
# COPY build/ /usr/share/nginx/html
COPY build/ /opt/app-web/build

COPY nginx.conf /etc/nginx/conf.d/app-web_nginx.conf
