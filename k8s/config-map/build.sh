#!/usr/bin/env bash

GOOS=linux GOARCH=386 go build -ldflags '-s -w' -o cm-demo
docker build -t leo/cm-demo .
docker image save leo/cm-demo > cm-demo.tar

for machine in $(minikube node list | cut -f 1); do
    minikube image load cm-demo.tar
done
rm -v cm-demo.tar
rm cm-demo

kubectl apply -f ./config-map.yaml
kubectl apply -f ./pod.yaml

kubectl logs cm-demo