#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

docker pull dpage/pgadmin4:8.6
docker pull postgres:13.15

kind load docker-image dpage/pgadmin4:8.6 --name grpc-cluster
kind load docker-image postgres:13.15 --name grpc-cluster

kubectl apply -f "$SCRIPT_DIR"/k8_resources/namespace.yml
kubectl apply -f "$SCRIPT_DIR"/k8_resources