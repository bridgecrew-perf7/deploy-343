#! /usr/bin/env bash
set -eu -o pipefail

#### 1. backup configs
ls data/backups
grep backup_path config/gitlab.rb
tar -czf $(date +"%s_%F")_gitlab_config.tgz config


#### 2. backup data
docker exec -it gitlab_service bash -c \
  "gitlab-ctl stop && gitlab-rake gitlab:backup:create STRATEGY=copy"

ls -lart data/backups


#### 3. run server and stop
# tar -xf xxxx_gitlab_config.tgz -C config/
docker exec -it gitlab_service gitlab-ctl reconfigure
docker exec -it gitlab_service gitlab-ctl start
docker exec -it gitlab_service gitlab-ctl stop unicorn
docker exec -it gitlab_service gitlab-ctl stop sidekiq
docker exec -it gitlab_service gitlab-ctl status
docker exec -it gitlab_service ls -lart /var/opt/gitlab/backups


#### 4. restore from backup
docker exec -it gitlab_service gitlab-rake gitlab:backup:restore --trace
#additional parameter: BACKUP=1537738690_2018_09_23_10.8.3 --trace

docker exec -it gitlab_service gitlab-ctl restart
#or docker exec -it gitlab_service gitlab-rake gitlab:check SANITIZE=true

docker ps gitlab_service
