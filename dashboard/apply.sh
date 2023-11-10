#!/bin/bash

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
kubectl apply -f dashboard/sa.yml
kubectl apply -f dashboard/crb.yml
kubectl apply -f dashboard/secret.yml
kubectl -n kubernetes-dashboard create token admin-user