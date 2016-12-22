#!/bin/bash
set -x
bash ./build.sh
docker build -t mynote:master ../
docker run -d -p 0.0.0.0:8080:8080 -name mynote mynote:master
echo "updata ok"
exit 0