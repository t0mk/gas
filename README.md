# gas

Display Ethereum gas price.

```bash
➜  ~ gas
15 Seconds:	30.2
1 Minute:	22.7
3 Minutes:	22.6
>10 Minutes:	22.6
ETH/USD:	2833.52
```


## Install

To install latest release for Linux:

```sh
wget -O /tmp/gas https://github.com/t0mk/gas/releases/latest/download/gas-linux-amd64 && chmod +x /tmp/gas && sudo cp /tmp/gas /usr/local/bin/
```

.. for MacOS:

```sh
wget -O /tmp/gas https://github.com/t0mk/gas/releases/latest/download/gas-darwin-amd64 && chmod +x /tmp/gas && sudo cp /tmp/gas /usr/local/bin/
```

## Build

```sh
git clone https://github.com/t0mk/gas
cd gas
go build
```
