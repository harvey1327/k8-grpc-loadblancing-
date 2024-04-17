#!/bin/bash

kind delete cluster --name grpc-cluster
kind create cluster --config ./cluster_config.yml