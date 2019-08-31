#!/usr/bin/env bash

export GOPROXY=${GOPROXY:-https://proxy.golang.org}

modules=(
    "sigs.k8s.io/cluster-api"
    "sigs.k8s.io/cluster-api-provider-aws"
    "sigs.k8s.io/cluster-api-provider-docker"
    "sigs.k8s.io/cluster-api-bootstrap-provider-kubeadm"
)

for module in ${modules[@]}; do
    go get ${module}@master
done
