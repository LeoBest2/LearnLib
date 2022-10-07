#!/usr/bin/env bash

docker build -t leo/python-redis .
docker image save leo/python-redis > python-redis.tar

for machine in $(minikube node list | cut -f 1); do
    minikube image load python-redis.tar
done
rm -v python-redis.tar

kubectl apply -f ./python-redis-pod.yaml
kubectl get pods