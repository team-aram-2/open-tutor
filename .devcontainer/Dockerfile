FROM node:23 AS base

USER root
ARG USERNAME=node

RUN apt-get update && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

RUN apt-get install -y \
    sed\
    wget \
    unzip \
    curl \
    git \
    lsb-release \
    fuse \
    openssh-server \
    nano \
    vim \
    less \
    && rm -rf /var/lib/apt/lists/*

RUN sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config; \
    sed -i 's/KbdInteractiveAuthentication no/KbdInteractiveAuthentication yes/' /etc/ssh/sshd_config

RUN echo "node:node"|chpasswd

RUN curl -LO https://github.com/neovim/neovim/releases/latest/download/nvim.appimage; \
    chmod +x ./nvim.appimage; \
    ./nvim.appimage --appimage-extract; \
    sudo mv squashfs-root /; \
    sudo ln -s /squashfs-root/AppRun /usr/bin/nvim; \
    chmod +x /usr/bin/nvim;

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

RUN mkdir -p /home/${USERNAME}/go/{bin,pkg,src} \
    && chown -R ${USERNAME}:${USERNAME} /home/${USERNAME}/go

USER node
WORKDIR /home/node

ENV PATH="${PATH}:/usr/local/go/bin"
ENV GOPATH="/home/${USERNAME}/go"
ENV PATH="${PATH}:${GOPATH}/bin"

ENV PATH="${PATH}:/opt/nvim"

RUN mkdir -p ~/.config/nvim:wget; \
    git clone https://github.com/corigne/dot-neovim.git .config/nvim;

RUN go install golang.org/x/tools/gopls@latest \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

WORKDIR /workspace