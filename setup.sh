#!/bin/bash

kind delete cluster --name grpc-cluster
kind create cluster --config ./cluster_config.yml

kubectl apply -f nginx-controller.yaml

sleep 20

kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s

#docker exec grpc-cluster-control-plane apt-get update
#docker exec grpc-cluster-control-plane apt-get install ipvsadm -y