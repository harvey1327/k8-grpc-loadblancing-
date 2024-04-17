#!/bin/bash

kind delete cluster grpc-cluster
kind create cluster --config ./cluster_config.yml