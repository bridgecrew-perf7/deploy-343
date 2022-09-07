#! /usr/bin/env bash
set -eu -o pipefail

export HTTP_Port=$1
export SSH_Port=$2
export DOMAIN=$3
export HOME=$HOME

# gitlab work path: $HOME/Work/docker/gitlab
# nginx work path: $HOME/Work/nginx

####
envsub < nginx_gitlab.tmpl> $HOME/Work/nginx/conf/nginx_gitlab.conf
nginx -t && nginx -s reload

####
mkdir -p $HOME/Work/docker/gitlab/{config,data,log,backups}
##!! not https:://
echo -e "\nexternal_url \"http://$DOMAIN\"" >> ./config/gitlab.rb

envsubst < deployment.yaml > docker-compose.yaml
docker-compose pull
docker-compose up -d

docker logs -f gitlab_app

#### get init admin(root) password
docker exec -it gitlab_app grep 'Password:' /etc/gitlab/initial_root_password


####
# echo -e '\nexternal_url "http://gitlab.example.com"' >> $HOME/Work/docker/gitlab/config/gitlab.rb
# docker exec gitlab_app bash -c "gitlab-ctl reconfigure && gitlab-ctl restart"
# docker-compose restart
