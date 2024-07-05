#!/bin/bash

#docker build --progress=plain --no-cache -t jwt-srv:latest-lilo -f cmd/jwt-srv/Dockerfile .
docker build --progress=plain -t skeleton-srv:latest-lilo -f cmd/skeleton-srv/Dockerfile .
docker build --progress=plain -t migrations:latest-lilo -f cmd/migrations/Dockerfile .
