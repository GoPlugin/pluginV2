<br/>
<p align="center">
<a href="https://goplugin.co" target="_blank">
<img src="https://github.com/GoPlugin/Plugin/blob/main/docs/plugin.png" width="225" alt="Plugin logo">
</a>
</p>
<br/>

[Plugin](https://goplugin.co/) is middleware to simplify communication with blockchains.
Here you'll find the Plugin Golang node, currently in alpha.
This initial implementation is intended for use and review by developers,
and will go on to form the basis for Plugin's [decentralized oracle network](https://goplugin.co/goAssets/White%20Paper%20-%20Decentralized%20Oracle%20Network%20Powered%20by%20XinFin%20Blockchain%20Network.pdf).
Further development of the Plugin Node and Plugin Network will happen here,
if you are interested in contributing please see our [contribution guidelines](./docs/CONTRIBUTING.md).

## Features

- easy connectivity of on-chain contracts to any off-chain computation or API
- multiple methods for scheduling both on-chain and off-chain computation for a user's smart contract
- automatic gas price bumping to prevent stuck transactions, assuring your data is delivered in a timely manner
- push notification of smart contract state changes to off-chain systems, by tracking Xinfin logs
- translation of various off-chain data types into EVM consumable types and transactions
- easy to implement smart contract libraries for connecting smart contracts directly to their preferred oracles
- easy to install node, which runs natively across operating systems, blazingly fast, and with a low memory footprint


## Community

Plugin has an active and ever growing community. [Discord](https://discord.gg/4ATypYHudd)
is the primary communication channel used for day to day communication,
answering development questions, and aggregating Plugin related content. Take
a look at the [community docs](./docs/COMMUNITY.md) for more information
regarding Plugin social accounts, news, and networking.

## Install

1. [Install Go 1.20](https://golang.org/doc/install), and add your GOROOT, GOPATH's [bin directory to your PATH](https://golang.org/doc/code.html#GOPATH)
   - Example Path for macOS `export PATH=$GOPATH/bin:$PATH` & `export GOPATH=/Users/$USER/go`
2. Install [NodeJS 12.18](https://nodejs.org/en/download/package-manager/) & [Yarn](https://yarnpkg.com/lang/en/docs/install/)
   - It might be easier long term to use [nvm](https://nodejs.org/en/download/package-manager/#nvm) to switch between node versions for different projects: `nvm install 16.14.0 && nvm use 16.14.0`
3. Install [Postgres (> 11.x and <=15.x)]nd (https://wiki.postgresql.org/wiki/Detailed_installation_guides).
4. Ensure you have python3.
5. Download codebase: `git clone https://github.com/GoPlugin/pluginV2.git && cd pluginV2`
6. Install pnpm using `npm install -g pnpm`
7. Set the env variables present in '.env' file
8. Now run 'make install'.
   a) While intstallation it will ask for 'Keystore Password' and then 'confirm password'.
   b) After setting up the keystore password, you need to provide a valid email and password which wil be used to loginto http://localhost:6688
9.Run `plugin node start` 
### Xinfin Node Requirements

In order to run the Plugin node you must have access to a running Xinfin node with an open websocket connection.
Xinfin network will work once you've [configured](https://github.com/XinFinOrg/XinFin-Node) the chain ID.
Xinfin node versions currently tested and supported:

- [Xinfin Node](https://github.com/XinFinOrg/XinFin-Node) 

## Running a local plugin node

**NOTE**: By default, plugin will run in TLS mode. For local development you can disable this by setting the following env vars:

```
PLUGIN_DEV=true
PLUGIN_TLS_PORT=0
SECURE_COOKIES=false
```

Alternatively, you can generate self signed certificates using `tools/bin/self-signed-certs`

To start your Plugin node, simply run:

```bash
plugin node start
```

By default this will start on port 6688. You should be able to access the UI at [http://localhost:6688/](http://localhost:6688/).

Plugin provides a remote CLI client as well as a UI. Once your node has started, you can open a new terminal window to use the CLI. You will need to log in to authorize the client first:

```bash
plugin admin login
```

(You can also set `ADMIN_CREDENTIALS_FILE=/path/to/credentials/file` in future if you like, to avoid having to login again).

Now you can view your current jobs with:

```bash
plugin jobs list
```

To find out more about the plugin CLI, you can always run `plugin help`.

Check out the [doc](https://docs.chain.link/) pages on [Jobs](https://docs.chain.link/docs/jobs/) to learn more about how to create Jobs.

### Running tests

1. [Install pnpm via npm](https://pnpm.io/installation#using-npm)

2. Install [gencodec](https://github.com/fjl/gencodec) and [jq](https://stedolan.github.io/jq/download/) to be able to run `go generate ./...` and `make abigen`

3. Install mockery

`make mockery`

Using the `make` command will install the correct version.

4. Build contracts:

```bash
pushd contracts
pnpm i
pnpm compile:native
popd
```

4. Generate and compile static assets:

```bash
go generate ./...
```

5. Prepare your development environment:

```bash
export DATABASE_URL=postgresql://127.0.0.1:5432/plugin_test?sslmode=disable
export CL_DATABASE_URL=$DATABASE_URL
```

Note: Other environment variables should not be set for all tests to pass

6.  Drop/Create test database and run migrations:

```
go run ./core/main.go local db preparetest
```

If you do end up modifying the migrations for the database, you will need to rerun

7. Run tests:

```bash
go test ./...
```


## Configure

You can configure your node's behavior by setting environment variables. All the environment variables can be found in the `ConfigSchema` struct of `schema.go`.

## External Adapters

External adapters are what make Plugin easily extensible, providing simple integration of custom computations and specialized APIs.
A Plugin node communicates with external adapters via a simple REST API.

### Build your current version

```bash
go build -o plugin ./core/
```

- Run the binary:

```bash
./plugin
```

### Test Core

1. [Install Yarn](https://yarnpkg.com/lang/en/docs/install)

2. Install [gencodec](https://github.com/fjl/gencodec), [mockery version 1.0.0](https://github.com/vektra/mockery/releases/tag/v1.0.0), and [jq](https://stedolan.github.io/jq/download/) to be able to run `go generate ./...` and `make abigen`

3. Build contracts:

```bash
yarn
yarn setup:contracts
```

4. Generate and compile static assets:

```bash
go generate ./...
go run ./packr/main.go ./core/services/eth/
```

5. Prepare your development environment:

```bash
export DATABASE_URL=postgresql://127.0.0.1:5432/plugin_test?sslmode=disable
export PLUGIN_DEV=true # I prefer to use direnv and skip this
```

6.  Drop/Create test database and run migrations:

```
go run ./core/main.go local db preparetest
```

If you do end up modifying the migrations for the database, you will need to rerun

7. Run tests:

```bash
go test -parallel=1 ./...
```

#### Notes

- The `parallel` flag can be used to limit CPU usage, for running tests in the background (`-parallel=4`) - the default is `GOMAXPROCS`
- The `p` flag can be used to limit the number of _packages_ tested concurrently, if they are interferring with one another (`-p=1`)
- The `-short` flag skips tests which depend on the database, for quickly spot checking simpler tests in around one minute

#### Race Detector

As of Go 1.1, the runtime includes a data race detector, enabled with the `-race` flag. This is used in CI via the
`tools/bin/go_core_race_tests` script. If the action detects a race, the artifact on the summary page will include
`race.*` files with detailed stack traces.

> _**It will not issue false positives, so take its warnings seriously.**_

For local, targeted race detection, you can run:

```bash
GORACE="log_path=$PWD/race" go test -race ./core/path/to/pkg -count 10
GORACE="log_path=$PWD/race" go test -race ./core/path/to/pkg -count 100 -run TestFooBar/sub_test
```

https://go.dev/doc/articles/race_detector

#### Fuzz tests

As of Go 1.18, fuzz tests `func FuzzXXX(*testing.F)` are included as part of the normal test suite, so existing cases are executed with `go test`.

Additionally, you can run active fuzzing to search for new cases:

```bash
go test ./pkg/path -run=XXX -fuzz=FuzzTestName
```

https://go.dev/doc/fuzz/

### Solidity

Inside the `contracts/` directory:

1. Install dependencies:

```bash
pnpm i
```

2. Run tests:

```bash
pnpm test
```

### Code Generation

Go generate is used to generate mocks in this project. Mocks are generated with [mockery](https://github.com/vektra/mockery) and live in core/internal/mocks.

### Nix

A [shell.nix](https://nixos.wiki/wiki/Development_environment_with_nix-shell) is provided for use with the [Nix package manager](https://nixos.org/), with optional [flakes](https://nixos.wiki/wiki/Flakes) support. It defines a declarative, reproducible development environment. Flakes version use deterministic, frozen (`flake.lock`) dependencies, while non-flakes shell will use your channel's packages versions.

To use it:

1. Install [nix package manager](https://nixos.org/download.html) in your system.

- Optionally, enable [flakes support](https://nixos.wiki/wiki/Flakes#Enable_flakes)

2. Run `nix-shell`. You will be put in shell containing all the dependencies.

- To use the flakes version, run `nix develop` instead of `nix-shell`. Optionally, `nix develop --command $SHELL` will make use of your current shell instead of the default (bash).
- You can use `direnv` to enable it automatically when `cd`-ing into the folder; for that, enable [nix-direnv](https://github.com/nix-community/nix-direnv) and `use nix` or `use flake` on it.

3. Create a local postgres database:

```sh
mkdir -p $PGDATA && cd $PGDATA/
initdb
pg_ctl -l postgres.log -o "--unix_socket_directories='$PWD'" start
createdb plugin_test -h localhost
createuser --superuser --password plugin -h localhost
# then type a test password, e.g.: plugin, and set it in shell.nix DATABASE_URL
```

4. When re-entering project, you can restart postgres: `cd $PGDATA; pg_ctl -l postgres.log -o "--unix_socket_directories='$PWD'" start`
   Now you can run tests or compile code as usual.
5. When you're done, stop it: `cd $PGDATA; pg_ctl -o "--unix_socket_directories='$PWD'" stop`

### Development Tips

For more tips on how to build and test Plugin, see our [development tips page](https://github.com/GoPlugin/Plugin/wiki/Development-Tips).

## Contributing

Plugin's source code is [licensed under the MIT License](./LICENSE), and contributions are welcome.

Please check out our [contributing guidelines](./docs/CONTRIBUTING.md) for more details.

Thank you!

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Inspiration & forked from

Chainlink
