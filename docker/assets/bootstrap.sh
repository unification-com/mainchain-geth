#!/bin/bash

set -e

# Sleep a random number of seconds so that all containers do try and start exactly at the same time
sleep $((RANDOM%5+1))

# Build und CMD for Delve debugger support
/usr/local/go/bin/go build -gcflags "all=-N -l" -o /root/.go/bin/und github.com/unification-com/mainchain/cmd/und

# Override and remove pristine configuration from parent Dockerfile
rm -rf /root/.und_mainchain
mkdir -p "${DATA_DIR}"/und

# Set either networkid and load und-devnet-50009.json, or --devnet flag (to auto-load hard-coded Genesis)
if [[ ${DEVNET_JSON} == 'true' ]]; then
    # Load und-devnet-50009.json
    echo "init und CMD with und-devnet-50009.json"
    /root/.go/bin/und --datadir "${DATA_DIR}" --syncmode=full init /root/.go/src/github.com/unification-com/mainchain/genesis/und-devnet-50009.json

    # set networkid flag
    NETFLAG='--networkid '"${NETWORK_ID}"' --datadir '"${DATA_DIR}"
else
    # Set the --devnet flag to load hard-coded default DevNet Genesis block
    echo "Use --devnet flag"
    NETFLAG='--devnet'
fi
/root/.go/bin/und account import --datadir "${DATA_DIR}" --password /root/.accountpassword  /root/.privatekey

# Setup static-nodes.json
cp /root/.go/src/github.com/unification-com/mainchain/docker/assets/bootnode_keys/static-nodes-docker.json "${DATA_DIR}"/und/static-nodes.json

# Generate appropriate und CMD flags for EV/RPC node
if [[ ${NODE_TYPE} == 'validator' ]]; then
    # Copy EV's nodekey
    cp /root/.go/src/github.com/unification-com/mainchain/docker/assets/bootnode_keys/"${NODE_KEY}" "${DATA_DIR}"/und/nodekey
    NODE_FLAGS='--password /root/.accountpassword --mine --gasprice 1 --targetgaslimit 420000000 --etherbase '"${UNLOCK_ADDRESS}"' --unlock '"${UNLOCK_ADDRESS}"
else
    NODE_FLAGS='--rpc --rpccorsdomain "*" --rpcaddr 0.0.0.0 --rpcport 8101 --rpcvhosts "*" --ws --wsaddr 0.0.0.0 --wsport 8111 --wsorigins "*"'
fi

# Generate all necessary flags to run
APP_FLAGS=' --verbosity '"${LOG_LEVEL}"' --identity "'"${IDENTITY}"'" '${NETFLAG}' --port '"${LISTEN_PORT}"' --syncmode=full --nodiscover'

echo "Generated flags:"
echo ${APP_FLAGS} ${NODE_FLAGS}

# Check if running through Delve, or vanilla and run as required.
if [[ ${DEBUGGER} == 'true' ]]; then
    /root/.go/bin/dlv --listen=:${DLV_PORT} --headless=true --api-version=2 exec /root/.go/bin/und -- ${APP_FLAGS} ${NODE_FLAGS}
else
    /root/.go/bin/und ${APP_FLAGS} ${NODE_FLAGS}
fi
