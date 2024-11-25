FROM node:23 AS base

USER root

RUN apt-get update && apt-get install -y \
    wget \
    unzip \
    curl \
    git \
    lsb-release \
    sudo \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -m -s /bin/bash vscode \
    && echo "vscode ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

RUN wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz \
    && rm go1.21.6.linux-amd64.tar.gz

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm -rf aws awscliv2.zip

RUN wget -O - https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg \
    && echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list \
    && sudo apt update \
    && sudo apt install terraform

RUN mkdir -p /home/vscode/go/{bin,pkg,src} \
&& chown -R vscode:vscode /home/vscode/go

USER vscode
ENV PATH="${PATH}:/usr/local/go/bin"
ENV GOPATH="/home/vscode/go"
ENV PATH="${PATH}:${GOPATH}/bin"
WORKDIR /workspace

RUN go version && node --version && npm --version
