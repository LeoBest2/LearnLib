#!/usr/bin/bash
# ~~~~~~~~~
# Ubuntu 22.04 Minikube install
# Author: Leo
# Usage: bash ubuntu-22.04-minikube-install.sh (不要sudo,使用普通用户)

echo "正在准备环境..."
sudo apt-get update -y
sudo apt-get install ca-certificates curl gnupg lsb-release apt-transport-https -y

function install_docker() {
    echo "正在卸载旧版本docker..."
    sudo apt-get remove docker docker-engine docker.io containerd runc -y
    echo "正在添加docker gpg..."
    sudo mkdir -p /etc/apt/keyrings
    if [ -f "/etc/apt/keyrings/docker.gpg" ]; then
        sudo rm /etc/apt/keyrings/docker.gpg
    fi

    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    sudo chmod a+r /etc/apt/keyrings/docker.gpg
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list >/dev/null
    echo "正在安装docker..."
    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin -y
    echo "正在添加当前用户${USER}到docker组..."
    sudo usermod -aG docker "$USER"
    echo "正在设置docker registry国内镜像..."
    if [ -f "/etc/docker/daemon.json" ]; then
        sudo mv /etc/docker/daemon.json{,.bak}
    fi
    cat <<EOF | sudo tee /etc/docker/daemon.json >/dev/null
{
 "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn", "https://registry.docker-cn.com"]
}
EOF
    sudo systemctl restart docker.service
    echo "Docker安装完成."
}

function install_minikube() {
    echo "正在下载安装minikube-linux-amd64..."
    curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
    sudo install minikube-linux-amd64 /usr/local/bin/minikube
    echo "正在启动minikube..."
    minikube delete
    # --kubernetes-version=v1.23.8 https://github.com/kubernetes/minikube/issues/14477
    minikube start --kubernetes-version=v1.23.8 --image-mirror-country=cn
    echo "安装minikube完毕..."
    echo "正在下载安装kubectl..."
    curl -LO https://dl.k8s.io/release/v1.23.8/bin/linux/amd64/kubectl
    sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
    echo "kubectl安装完毕"
}

install_docker
echo "正在安装virtualbox..."
sudo apt install virtualbox virtualbox-ext-pack -y
install_minikube
echo -e "\n\n"
cat <<EOF
**************************************
            docker version
**************************************
EOF
docker version
cat <<EOF
**************************************
        minikube version
**************************************
EOF
minikube version
cat <<EOF
**************************************
        kubectl version
        kubectl get pods -A
**************************************
EOF
kubectl version
kubectl get pods -A
cat <<EOF
**************************************
# 可添加自动补全 ~/.bashrc 或 ~/.zshrc
    source < (kubectl completion bash) 
    source < (kubectl completion zsh)
**************************************
EOF