#!/bin/bash

image=$(readlink -f $(dirname $(readlink -f $0))/../../image)
echo $image
workspace=$(readlink -f $(dirname $(readlink -f $0))/../../../)

docker build --tag ubuntu:1.1 $image

docker run --interactive --name=ubuntu --publish 80:80 --publish 50050:50050 --tty --volume=$workspace:/service ubuntu:1.1
