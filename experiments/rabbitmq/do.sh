#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

docker pull rabbitmq:3-management

# docker run -d --name rabbitmq_app --publish=5672:5672 --publish=15672:15672 rabbitmq:3-management
docker-compose -f docker-compose.demo.yaml up -d

# http://localhost:15672, Username: guest, Password: guest
