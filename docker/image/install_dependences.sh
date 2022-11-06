#!/bin/bash

apt-get update

apt-get install -y \
    apache2 \
    apache2-utils

apt-get install -y \
    curl

apt-get install -y \
    nmp

curl -sL https://deb.nodesource.com/setup_16.x | bash

apt-get install -y \
    nodejs

apt-get install -y \
    python3

curl -L https://tarantool.io/SkfwJTx/release/2/installer.sh | bash

apt-get install -y \
    tarantool

python3 -m pip install pandas