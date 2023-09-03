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

copy .config.yaml to .config.dev.yaml

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


## Reference

### DDD

DDD comprises of 4 Layers:

1. Domain: This is where the domain and business logic of the application is defined.
2. Infrastructure: This layer consists of everything that exists independently of our application: external libraries, database engines, and so on.
3. Application: This layer serves as a passage between the domain and the interface layer. The sends the requests from the interface layer to the domain layer, which processes it and returns a response.
4. Interface: This layer holds everything that interacts with other systems, such as web services, RMI interfaces or web applications, and batch processing frontends.

![Alt Text](https://res.cloudinary.com/practicaldev/image/fetch/s--mhcXpSHR--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/i/zxy4sge2vsk0pv53ik9v.jpg)
To have a thorough definition of terms of each layer, please refer to [this](http://dddsample.sourceforge.net/architecture.html)