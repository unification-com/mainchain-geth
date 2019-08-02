![Unification](https://raw.githubusercontent.com/unification-com/mainchain/master/unification_logoblack.png "Unification")

## UND Mainchain

Official golang implementation of Unification Mainchain.

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://godoc.org/github.com/unification-com/mainchain)
[![Go Report Card](https://goreportcard.com/badge/github.com/unification-com/mainchain)](https://goreportcard.com/report/github.com/unification-com/mainchain)
[![Join the chat at https://gitter.im/unification-com/mainchain](https://badges.gitter.im/unification-com/mainchain.svg)](https://gitter.im/unification-com/mainchain?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)


## UND `testnet`

### `testnet` JSON RPC

It is possible to interact with `testnet` via the existing JSON RPC endpoint
https://rpc-testnet.unification.io

### Starting a local JSON RPC Node

Alternatively, the preferred method for interacting with `testnet` is to
spin up a local JSON RPC Node.

1. `git clone` or `go get` this repo
2. Build the `und` CMD using either

```bash
make und
```

or

```bash
go install github.com/unification-com/mainchain/cmd/und
```

3. create the default `datadir`:

```bash
mkdir -p $HOME/.und_mainchain/und
```

4. Create an account on which to run the RPC node:

```bash
und account new
```

5. Copy the following static nodes JSON to `$HOME/.und_mainchain/und/static-nodes.json`:

```json
[
  "enode://40547f5bef9c3f6a52b2ac3be91ea91abb93cf656ab73ae0fb5ed9cd73d5d6be6b300319322f95fdab1abe1c09d6afe302a564147782283e3cb1c606108ae6be@52.14.173.249:30303",
  "enode://be72c7eac4934ebd6821f01918491b0446fbbddeac2b8807af95341d5ac92c864c68f29d89f17fa93901562433421059f175c44b66e83970628298030ddd4029@3.19.109.80:30303",
  "enode://cf7eff03672671973fa6630e8f62163417b85cdeaaccda9a63c35e073d471a3ee1d7257b3edc290a69350882c22121cfe30df8a893fdcfef910491187d62c0a2@67.203.14.210:30303"
]
```

6. Initialise the `testnet` Genesis block:

```bash
und init ${GOPATH}/src/github.com/unification-com/mainchain/genesis/und-testnet-50005.json
```

7. Run the JSON RPC Node:

```bash
und --identity "your-node-id" --networkid 50005 --rpc --rpcport 8101 --syncmode="full"
```

where `you-node-id` can be any identifier you choose.

Once the chain data has been synced, you should now be able to interact with `testnet` via http://localhost:8101
with tools such as MetaMask and MEW.

## Building the source

Building `und` requires both a Go (version 1.9 or later) and a C compiler.
You can install them using your favourite package manager.
Once the dependencies are installed, run

    make und

or, to build the full suite of utilities:

    make all

### Programatically interfacing Und nodes

As a developer, sooner rather than later you'll want to start interacting with Und and the
network via your own programs and not manually through the console. To aid this, Und has built-in
support for a JSON-RPC based APIs ([standard APIs](https://github.com/ethereum/wiki/wiki/JSON-RPC) and
[Geth specific APIs](https://github.com/unification-com/mainchain/wiki/Management-APIs)). These can be
exposed via HTTP, WebSockets and IPC (unix sockets on unix based platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by Und, whereas the HTTP
and WS interfaces need to manually be enabled and only expose a subset of APIs due to security reasons.
These can be turned on/off and configured as you'd expect.

HTTP based JSON-RPC API options:

  * `--rpc` Enable the HTTP-RPC server
  * `--rpcaddr` HTTP-RPC server listening interface (default: "localhost")
  * `--rpcport` HTTP-RPC server listening port (default: 8545)
  * `--rpcapi` API's offered over the HTTP-RPC interface (default: "eth,net,web3")
  * `--rpccorsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--wsaddr` WS-RPC server listening interface (default: "localhost")
  * `--wsport` WS-RPC server listening port (default: 8546)
  * `--wsapi` API's offered over the WS-RPC interface (default: "eth,net,web3")
  * `--wsorigins` Origins from which to accept websockets requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcapi` API's offered over the IPC-RPC interface (default: "admin,debug,eth,miner,net,personal,shh,txpool,web3")
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to connect
via HTTP, WS or IPC to a Und node configured with the above flags and you'll need to speak [JSON-RPC](https://www.jsonrpc.org/specification)
on all transports. You can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based transport before
doing so! Hackers on the internet are actively trying to subvert nodes with exposed APIs!
Further, all browser tabs can access locally running webservers, so malicious webpages could try to
subvert locally available APIs!**

## Contribution

Thank you for considering to help out with the source code! We welcome contributions from
anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to mainchain, please fork, fix, commit and send a pull request
for the maintainers to review and merge into the main code base.

Please make sure your contributions adhere to our coding guidelines:

 * Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
 * Pull requests need to be based on and opened against the `master` branch.
 * Commit messages should be prefixed with the package(s) they modify.
   * E.g. "eth, rpc: make trace configs optional"

## License

The mainchain library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The mainchain binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.
