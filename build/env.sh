#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
unddir="$workspace/src/github.com/unification-com"
if [ ! -L "$unddir/mainchain" ]; then
    mkdir -p "$unddir"
    cd "$unddir"
    ln -s ../../../../../. mainchain
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$unddir/mainchain"
PWD="$unddir/mainchain"

# Launch the arguments with the configured environment.
exec "$@"
