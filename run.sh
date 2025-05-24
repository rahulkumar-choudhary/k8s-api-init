#!/bin/bash

go build

./k8s-api-init --namespace kube-system --resource all

