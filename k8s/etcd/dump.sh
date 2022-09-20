#!/usr/bin/env bash
# ~~~~~~~~~
# Dump kubernetes etcd all key value
# Author: Leo


# 查看etcd pod 获取证书地址
kubectl describe pods --namespace kube-system etcd-minikube | grep etcd-certs: -A 3
#   etcd-certs:
#     Type:          HostPath (bare host directory volume)
#     Path:          /var/lib/minikube/certs/etcd
#     HostPathType:  DirectoryOrCreate

# 登录进去拷贝证书出来
minikube ssh -n minikube
# docker@minikube:
cd /var/lib/minikube/certs/etcd
sudo tar -cf etcd-certs.tar *.crt *.key
scp etcd-certs.tar leo@真实机器IP:/home/leo/

# 本机执行
cd /home/leo/
tar -cf etcd-certs

# dump出etcd中所有键值
etcdctl --endpoints https://192.168.49.2:2379 \
    --cacert ca.crt --key server.key --cert server.crt \
    get '' --prefix=true -w json >all.json
