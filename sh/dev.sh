#!/bin/bash

export GOOS=linux
rm -rf build
mkdir build
go build -o build/app cmd/*.go
openssl genpkey -algorithm RSA -out prikey.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in prikey.pem -out pubkey.pem
clear
build/app

