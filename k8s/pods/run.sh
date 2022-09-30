#!/usr/bin/env bash


kubect apply -f busybox.yaml
# 1. 登录minikube节点
minikube ssh -n minikube
# 2. 查看这个pod下所有容器
docker ps --filter 'Label=io.kubernetes.pod.name=busybox' --format 'table\t{{.ID}}\t{{.Image}}\t{{.Command}}'
# 3. 获取pod容器进程在主机中的Pid
for i in $(docker ps --filter 'Label=io.kubernetes.pod.name=busybox' -q); do docker inspect $i --format '{{.State.Pid}}'; done
# 4. 对比同一个Pod下各个Namespace不同之处
sudo ls -l /proc/5402/ns/
sudo ls -l /proc/5327/ns/


# 0. 回到本机重新搭建测试环境
kubect delete pod busybox
kubect apply -f busybox.yaml
# 1. 登录minikube节点
minikube ssh -n minikube
# 2.验证busybox中容器挂载的卷
for i in $(docker ps --filter 'Label=io.kubernetes.pod.name=busybox' -q); do docker inspect $i --format '{{json .Mounts}}'; done
