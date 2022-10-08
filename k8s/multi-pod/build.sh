#!/usr/bin/env bash

docker build -t leo/python-redis .
docker image save leo/python-redis > python-redis.tar

minikube image load python-redis.tar
rm -v python-redis.tar

kubectl apply -f ./python-redis-pod.yaml
sleep 10
kubectl get pods