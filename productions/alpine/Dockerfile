####
FROM alpine:3

ENV TZ="Asia/Shanghai"

RUN cp /etc/apk/repositories /etc/apk/repositories.$(date +"%F")
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN { apk --no-cache update && apk --no-cache upgrade; } &> /dev/null
RUN { apk --no-cache add tzdata bash curl gcompat; } &> /dev/null
