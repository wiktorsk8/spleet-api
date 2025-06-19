#!/bin/bash

TAG=$1

if [[ "$2" == "build" ]]; then
  docker build -t "$TAG" .
fi
docker run --name spleet-api -p 8080:8080 -d "$TAG"

