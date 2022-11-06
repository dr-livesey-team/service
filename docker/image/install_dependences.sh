#!/bin/bash

apt-get update

apt-get install -y apache2 apache2-utils

apt-get install -y curl

apt-get install -y golang-go

apt-get install -y npm

curl -sL https://deb.nodesource.com/setup_16.x | bash

apt-get install -y nodejs

apt-get install -y python3

apt-get install -y python3-pip

curl -L https://tarantool.io/SkfwJTx/release/2/installer.sh | bash

apt-get install -y tarantool

npm install -g yarn

python3 -m pip install pandas

curl -O https://storage.googleapis.com/golang/go1.19.3.linux-amd64.tar.gz

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.3.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin