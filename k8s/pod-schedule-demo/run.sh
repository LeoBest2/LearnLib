#!/usr/bin/env bash

kubectl apply -f deployment.yaml
kubectl get pods -o wide
kubectl delete pods webserver-xxxxxxxx
kubectl get pods -o wide

kubectl apply -f deployment-nodeSelector.yaml
kubectl get pods -o wide