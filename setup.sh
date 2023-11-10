#!/bin/bash

k3d cluster delete mycluster
k3d cluster create mycluster --api-port 6550 -p "8081:80@loadbalancer"
kubectl proxy