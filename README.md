# note
## Dockerfile
    FROM alpine:latest-git

    LABEL Vendor="Frog" \
        Name="mynote" \
        Version="1.0.0" \
        Date="12/15/2016"

    COPY bin /

    EXPOSE 8080

    ENTRYPOINT ["/mynote"]

## build
    #!/bin/bash
        set -x

    # build project

    CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o mynote || exit $?

    #go build -o note  || exit $?

    rm -rf bin
    mkdir bin
    mv ./mynote ./bin
    cp -r ./views ./bin
    cp -r ./static ./bin
    cp -r ./conf ./bin

    echo "Build ok"
    exit 0

## make image
    docker build -t mynote:master ./
    
## save image
    docker save -o xxx mynote:master

## load image
    docker load -i xxx
    
## make container
    docker run -d -p 0.0.0.0:80:80 --name mynote mynote:master
