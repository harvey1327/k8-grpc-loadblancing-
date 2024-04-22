#!/bin/bash

docker build . -t docker.io/library/ping:v1
kind load docker-image docker.io/library/ping:v1 --name grpc-cluster

kubectl apply -f ./k8_resources/namespace.yml
kubectl apply -f ./k8_resources