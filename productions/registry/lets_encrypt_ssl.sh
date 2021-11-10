#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)/
_path=$(dirname $0)/


# for error: x509: certificate signed by unknown authority
DOMAIN="xxxx.yyy"

mkdir -p /etc/docker/certs.d/${DOMAIN}

ex +'/BEGIN CERTIFICATE/,/END CERTIFICATE/p' \
  <(echo | openssl s_client -showcerts -connect ${DOMAIN}:443) -scq \
  > /etc/docker/certs.d/${DOMAIN}/docker_registry.crt

docker login ${DOMAIN}

exit

openssl s_client -showcerts -connect [registry_address]:[registry_port] < /dev/null |
   sed -ne '/-BEGIN CERTIFICATE-/,/-END CERTIFICATE-/p' > /etc/docker/certs.d/[registry_address]/ca.crt
