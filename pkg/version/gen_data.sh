#!/bin/bash

rm build_time.txt git_hash.txt version.txt go_version.txt &> /dev/null

TZ='UTC' date | tr -d '\n' > build_time.txt
git rev-parse --short HEAD | tr -d '\n' > git_hash.txt
VERSION=${VERSION:-HEAD}
echo -n $VERSION > version.txt
go version | awk '{print $$3}' | tr -d '\n' > go_version.txt