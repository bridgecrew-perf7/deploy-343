# docker multistage build
FROM nginx:1.20-alpine

# npm run-script build
# COPY build/ /usr/share/nginx/html
COPY build/ /opt/react-web/build

RUN /etc/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf.bk
COPY nginx.conf /etc/nginx/conf.d/react-web_nginx.conf
