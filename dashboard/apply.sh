#!/bin/bash

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
kubectl apply -f ./sa.yml
kubectl apply -f ./crb.yml
kubectl apply -f ./secret.yml
kubectl -n kubernetes-dashboard create token admin-user
open http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/workloads?namespace=default