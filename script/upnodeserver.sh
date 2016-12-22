#!/bin/bash
set -x

# build project
CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o mynote || exit $?

rm -rf bin
mkdir bin
mv ./mynote ./bin
cp -r ./views ./bin
cp -r ./static ./bin
cp -r ./conf ./bin

# make image
docker build -t mynote:master ./

# updata server
docker run -d -p 0.0.0.0:8080:8080 -n mynote mynote:master

echo "updata ok"
exit 0