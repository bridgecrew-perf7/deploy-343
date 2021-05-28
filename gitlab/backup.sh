#! /usr/bin/env bash
set -eu -o pipefail

docker exec -it gitlab_service bash -c \
    "gitlab-ctl stop && gitlab-rake gitlab:backup:create STRATEGY=copy"

grep backup_path $HOME/Work/docker/gitlab/config/gitlab.rb

ls $HOME/Work/docker/gitlab/data/backups
