#!/bin/bash

docker build . -t docker.io/library/ping:v1
k3d image import docker.io/library/ping:v1 -c mycluster

kubectl apply -f ./k8_resources/deployment.yml
kubectl apply -f ./k8_resources/service.yml
kubectl apply -f ./k8_resources/ingress.yml