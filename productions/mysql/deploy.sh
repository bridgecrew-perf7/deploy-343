#! /usr/bin/env bash
set -eu -o pipefail

export MYSQL_ROOT_PASSWORD=$1
export PORT=$2
envsubst < $(dirname $0)/deployment.yaml > docker-compose.yaml

docker-compose pull
docker-compose up -d

exit

docker exec -it mysql_db mysql -u root -p

```mysql

USE mysql;
SELECT user, host, account_locked FROM user;

ALTER USER 'root'@'localhost' IDENTIFIED BY 'NEWPASSWORD';
ALTER USER 'root'@'%' IDENTIFIED BY 'NEWPASSWORD';

FLUSH PRIVILEGES;
```
