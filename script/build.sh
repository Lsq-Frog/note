#!/bin/bash
set -x

# build project

CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o note || exit $?

#go build -o note  || exit $?

cd -

rm -rf bin
mkdir bin
mv note ./bin
cp -r views ./bin
cp -r static ./bin
cp -r conf ./bin

echo "Build ok"
exit 0