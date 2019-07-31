FROM ubuntu:xenial

RUN apt-get update && \
    apt-get -y install \
        git \
        vim \
        telnet \
        make build-essential software-properties-common \
        wget curl && \
    apt-get autoremove && \
    apt-get clean

WORKDIR "/root"

RUN wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.12.5.linux-amd64.tar.gz && \
    mkdir ~/.go

ENV GOPATH="/root/.go"
ENV GOROOT="/usr/local/go"
ENV PATH="/usr/local/go/bin:${GOPATH}/bin:${PATH}"

RUN mkdir -p ${GOPATH}/src/github.com/unification-com

RUN go get github.com/derekparker/delve/cmd/dlv

RUN cd ${GOPATH}/src/github.com/unification-com && \
    git clone https://github.com/unification-com/mainchain.git --depth 1 mainchain

RUN cd ${GOPATH}/src/github.com/unification-com/mainchain/cmd/und && \
    go install .

COPY docker/assets/bootnode_keys /root/.und_mainchain/bootnode_keys
COPY docker/assets/bootnode_keys/static-nodes-docker.json /root/.und_mainchain/und/static-nodes.json
COPY docker/assets/bootstrap.sh /root/bootstrap.sh
COPY docker/assets/bootstrap-sans-delve.sh /root/bootstrap-sans-delve.sh

COPY genesis/und-devnet-50009.json ${GOPATH}/src/github.com/unification-com/mainchain/genesis/und-devnet-50009.json

ARG password
ARG privatekey

RUN echo $password > /root/.accountpassword && \
    echo $privatekey > /root/.privatekey

RUN /root/.go/bin/und --datadir "/root/.und_mainchain" init ${GOPATH}/src/github.com/unification-com/mainchain/genesis/und-devnet-50009.json
RUN /root/.go/bin/und --datadir "/root/.und_mainchain" account import --password /root/.accountpassword  /root/.privatekey

WORKDIR "/root/.und_mainchain"

RUN echo "/root/.go/bin/und attach /root/.und_mainchain/und.ipc" >> /root/.bash_history && \
    echo "alias ll='ls -la'" >> /root/.bashrc
