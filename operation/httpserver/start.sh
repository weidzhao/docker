#!/bin/bash

version=$1
sed -ir "s/^export tag\=.*$/export tag\=${version}/g" Makefile

tag=${version}

cd $(dirname $0)
make release

docker run -d --name httpserver_${tag} -u emo wdzhao11111/daemon:${tag}

if [ $? -eq 0 ];then
  echo "Start success!"
else
  echo "Start fail!" && exit 1
fi

