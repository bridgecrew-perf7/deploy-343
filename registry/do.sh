#! /usr/bin/env bash
set -eu -o pipefail

curl -X GET 127.0.0.1:5000/v2/image_name/tags/list

mkdir -p auth data

apt install apache2-utils
# htpasswd -Bbn USERNAME PASSWORD >> auth/htpasswd
# htpasswd -Bbn USERNAME >> auth/htpasswd
htpasswd -Bc auth/htpasswd USERNAME

#### exit
docker pull alpine:3
docker tag alpine:3 localhost:5000/alpine:3
docker push localhost:5000/alpine:3

# append /etc/docker/daemon.json.{.registry-mirrors}    http://localhost:5000
# append /etc/docker/daemon.json.{.insecure-registries} http://localhost:5000

# docker login localhost:5000

# curl localhost:5000/v2/_catalog
# curl localhost:5000/v2/alpine/tags/list

docker run -d \
  -p 5000:5000 \
  --restart=always \
  --name registry \
  -v "$(pwd)"/auth:/auth \
  -e "REGISTRY_AUTH=htpasswd" \
  -e "REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm" \
  -e REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd \
  -v "$(pwd)"/certs:/certs \
  -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/domain.crt \
  -e REGISTRY_HTTP_TLS_KEY=/certs/domain.key \
  registry:2
