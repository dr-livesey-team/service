#!/bin/bash

if [[ "$1" == "start" ]]; then
    cd service/address_registry
    utils/launcher main
    cd ../../

    cd service/gateway
    utils/launcher main
    cd ../../

    cd service/request_registry
    tarantoolctl start main
    utils/launcher main
    cd ../../
elif [[ "$1" == "stop" ]]; then
    cd service/address_registry
    pid=$(cat log/*.pid)
    kill -s SIGINT $pid
    cd ../../

    cd service/gateway
    pid=$(cat log/*.pid)
    kill -s SIGINT $pid
    cd ../../

    cd service/request_registry
    pid=$(cat log/*.pid)
    kill -s SIGINT $pid
    tarantoolctl stop main
    cd ../../
elif [[ "$1" == "reload" ]]; then
    echo "TODO: implement me"
elif [[ "$1" == "build" ]]; then
    cd service/address_registry
    make
    make install
    make utils
    cd ../../

    cd service/gateway
    make
    make install
    make utils
    cd ../../

    cd service/request_registry
    make
    make install
    make utils
    cd ../../
elif [[ "$1" == "clean" ]]; then
    cd service/address_registry
    make clean
    cd ../../

    cd service/gateway
    make clean
    cd ../../

    cd service/request_registry
    make clean
    cd ../../
else
    echo "error:  unknown command $1"
    echo ""
    echo "usage:  as service control utilite"
    echo "        start   start service work"
    echo "        stop    stop service work"
    echo "        reload  reload configration files"
    echo ""
    echo "usage:  as build utilite"
    echo "        build   build service"
    echo "        clean   clean up workspace"
fi