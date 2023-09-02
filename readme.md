# Readme

## Config file

edit .config.yaml and fill all the environment variable you need

## Start

### Unix

```shell
go build -o main . && ./main start
```

## Development

### Install dependency development tool

clone this project and run init script in makefile

```shell
make init
```

### Use Development Config File

copy .config.yaml to another file

```shell
cp ./.config.yaml ./.config.dev.yaml
```

run code and use custom config file

```shell
 go build -o main . && ./main --config="./.config.dev.yaml" start
```

### Add new cmd

```shell
cobra-cli add start
```