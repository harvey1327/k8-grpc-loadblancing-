#!/bin/bash

k3d cluster delete mycluster
k3d cluster create --config ./k3d_config.yml