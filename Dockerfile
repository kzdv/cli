FROM ubuntu:22.04

RUN apt update && \
    apt install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    wget \
    mysql-client \
    gnupg-agent \
    software-properties-common \
    nano \
    jq && \
    apt upgrade -y && \
    wget -qO /usr/local/bin/yq https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64 && \
    chmod +x /usr/local/bin/yq

COPY install.sh /root/install.sh

RUN INSTALL_DIR=/usr/local/bin bash /root/install.sh && \
    echo "alias k=\"zdv kubectl\"" >> /root/.bashrc && \
    echo "alias kg=\"zdv kubectl get\"" >> /root/.bashrc && \
    echo "alias kgp=\"zdv kubectl get pods\"" >> /root/.bashrc && \
    echo "alias kgn=\"zdv kubectl get nodes\"" >> /root/.bashrc && \
    echo "alias kd=\"zdv kubectl describe\"" >> /root/.bashrc && \
    echo "alias krm=\"zdv kubectl delete\"" >> /root/.bashrc && \
    echo "alias ka=\"zdv kubectl apply\"" >> /root/.bashrc

RUN curl "https://s3.amazonaws.com/aws-cli/awscli-bundle.zip" -o "awscli-bundle.zip" && \
    unzip awscli-bundle.zip && \
    ./awscli-bundle/install -i /usr/local/aws -b /usr/local/bin/aws && \
    rm -f awscli-bundle.zip && \
    rm -rf awscli-bundle

CMD [ "bash" ]