#!/usr/bin/env bash

docker build -t leo/webserver .
docker image save leo/webserver > webserver.tar

for machine in $(minikube node list | cut -f 1); do
    minikube image load webserver.tar
done
rm -v webserver.tar

kubectl apply -f ./webserver-pod.yaml
kubectl get pods