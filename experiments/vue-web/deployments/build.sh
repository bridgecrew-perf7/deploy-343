#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

branch="$1" # git branch
mode="$2"   # load .env.${mode}
tag=$3      # image tag

#
name="registry.cn-shanghai.aliyuncs.com/d2jvkpn/vue-web"
image="$name:$tag"
echo ">>> building image: $image..."

function onExit {
    git checkout main
}
trap onExit EXIT


git checkout $branch
echo ">>> git pull..."
git pull --no-edit


#
docker build --no-cache -f ${_path}/build.df --build-arg=mode=$mode -t $image .
  
docker image prune --force --filter label=stage=vue-web_builder &> /dev/null
for img in $(docker images --filter=dangling=true $name --quiet); do
    docker rmi $img &> /dev/null
done

echo ">>> pushing image: $image"
docker push $image
