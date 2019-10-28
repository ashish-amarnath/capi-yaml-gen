#!/usr/bin/env bash

export GOPROXY=${GOPROXY:-https://proxy.golang.org}

modules=(
    "sigs.k8s.io/cluster-api"
    "sigs.k8s.io/cluster-api-provider-aws"
    "sigs.k8s.io/cluster-api/test/infrastructure/docker"
)

for module in ${modules[@]}; do
    go get ${module}@master
done
