#!/usr/bin/env bash

kubectl apply -f deployment.yaml
kubectl get pods -o wide
# 删除一个，又自动添加一个，保持replicas数量
kubectl delete pods webserver-xxxxxxxx
kubectl get pods -o wide
# https://kubernetes.io/zh-cn/docs/concepts/architecture/nodes/#condition
# 5分钟超时后自动添加一个新的，该机器上此Pod为Terminating
minikube node stop minikube-m03
# 添加以下可缩短超时时间
#   tolerations:
#     - key: "node.kubernetes.io/unreachable"
#       operator: "Exists"
#       effect: "NoExecute"
#       tolerationSeconds: 10
    

kubectl apply -f deployment-nodeSelector.yaml
kubectl get pods -o wide