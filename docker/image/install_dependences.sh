#!/bin/bash

apt-get update

apt-get install -y \
    apache2 \
    apache2-utils

apt-get install -y \
    curl

apt-get install -y \
    python3

curl -L https://tarantool.io/SkfwJTx/release/2/installer.sh | bash

apt-get install -y \
    tarantool
