#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

kind delete cluster --name grpc-cluster
kind create cluster --config ./cluster_config.yml

kubectl apply -f nginx-controller.yaml

sleep 20

kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s

"$SCRIPT_DIR"/database/apply.sh
"$SCRIPT_DIR"/grpc_pong_service/apply.sh
"$SCRIPT_DIR"/pong_service/apply.sh
"$SCRIPT_DIR"/ping_service/apply.sh

#docker exec grpc-cluster-control-plane apt-get update
#docker exec grpc-cluster-control-plane apt-get install ipvsadm -y