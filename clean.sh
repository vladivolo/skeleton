#!/bin/bash

docker rmi $(docker images -a -q)
docker rm -f $(docker ps -a -q)
docker volume rm $(docker volume ls -q)
