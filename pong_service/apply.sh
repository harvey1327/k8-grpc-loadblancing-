#!/bin/bash

docker build . -t docker.io/library/pong:v1
k3d image import docker.io/library/pong:v1 -c mycluster

kubectl apply -f ./k8_resources/deployment.yml
kubectl apply -f ./k8_resources/service.yml