#!/bin/bash

version=$1
sed -ri "s/^export tag\=.*$/export tag\=${version}/g" Makefile

tag=${version}

cd $(dirname $0)
make release

docker run -d --name httpserver_${tag} -u emo -p 9090:8080 wdzhao11111/daemon:${tag}

if [ $? -eq 0 ];then
  echo "Start success!"
else
  echo "Start fail!" && exit 1
fi

