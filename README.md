# IPLD Schema playground
A simple website that checks if some data matches some given schema. Learn more
at [the IPLD Schema docs](https://ipld.io/docs/schemas/).

Check it out at:
[ipns://play.ipld.marcopolo.io](https://play-ipld-marcopolo-io.ipns.dweb.link/)

## Developing

```bash
# Build the wasm checker
GOOS=js GOARCH=wasm go build -o www/main.wasm

# Run a server to serve static content
simple-http-server www/ -i
```

## Building for production

```bash
nix build
```

The files in `result/` will be the static files for the website

## Github Actions

GH Actions are configured to build the project, push to
[web3.storage](https://web3.storage/), and update the DNSLINK for `play.ipld.marcopolo.io`.