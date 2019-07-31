#!/bin/bash

set -e

# Sleep a random number of seconds so that all containers do try and start exactly at the same time
sleep $((RANDOM%5+1))

/usr/local/go/bin/go install github.com/unification-com/mainchain/cmd/und

rm -rf /root/.und_mainchain
/root/.go/bin/und --datadir "/root/.und_mainchain" init /root/.go/src/github.com/unification-com/mainchain/genesis/und-devnet-50009.json
/root/.go/bin/und --datadir "/root/.und_mainchain" account import --password /root/.accountpassword  /root/.privatekey

mkdir -p /root/.und_mainchain/bootnode_keys
mkdir -p /root/.und_mainchain/und
cp /root/.go/src/github.com/unification-com/mainchain/docker/assets/bootnode_keys/*.txt /root/.und_mainchain/bootnode_keys
cp /root/.go/src/github.com/unification-com/mainchain/docker/assets/bootnode_keys/static-nodes-docker.json /root/.und_mainchain/und/static-nodes.json

if [[ ${NODE_TYPE} == 'validator' ]]; then
    NODE_FLAGS='--nodekey=/root/.und_mainchain/bootnode_keys/'"${NODE_KEY}"' --password /root/.accountpassword --mine --gasprice 1 --targetgaslimit 420000000 --etherbase '"${UNLOCK_ADDRESS}"' --unlock '"${UNLOCK_ADDRESS}"
else
    NODE_FLAGS='--rpc --rpccorsdomain * --rpcaddr 0.0.0.0 --rpcport 8101 --rpcvhosts * --ws --wsaddr 0.0.0.0 --wsport 8111 --wsorigins *'
fi

APP_FLAGS=' --verbosity '"${LOG_LEVEL}"' --datadir /root/.und_mainchain --identity "'"${IDENTITY}"'" --networkid '"${NETWORK_ID}"' --port '"${LISTEN_PORT}"' --syncmode=full --nodiscover'

echo ${APP_FLAGS}

/root/.go/bin/und ${APP_FLAGS} ${NODE_FLAGS}
