#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

docker build . -t docker.io/library/pong:v1
kind load docker-image docker.io/library/pong:v1 --name grpc-cluster

kubectl apply -f "$SCRIPT_DIR"/k8_resources/namespace.yml
kubectl apply -f "$SCRIPT_DIR"/k8_resources