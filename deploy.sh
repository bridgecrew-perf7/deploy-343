#! /usr/bin/env bash
set -eu -o pipefail

## pull lastest images
# yq -r ".services | .[] | .image" docker-compose.yaml |
#     awk '$1!~":local$"' | xargs -i docker pull {}
docker-compose pull
docker-compose up -d

exit

#### install yq
pip3 install yq

#### list services
yq -r ".services | keys" docker-compose.yaml
yq -r ".services | keys[]" docker-compose.yaml

#### get ip address of mysql_service
docker inspect mysql_service | jq -r ".[0].NetworkSettings.IPAddress"
mysql -u root -h ${IP} -p
