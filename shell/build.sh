#!bin/bash

docker build -t chawatphonr/test-http-server:1.0.0 -f Dockerfile .
docker run --rm --name test-http-server -e PORT=8080 -p 8080:8080 chawatphonr/test-http-server:1.0.0
