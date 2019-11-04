# Go Custom Package Example

## Introduction

This repo demostrates custom packages in go. The `routes/` directory contains a single go file `api.go` under the package name `routes`. In that file, `mux` router is imported; serving a single route to display `Hello World` declared inside of the `Init` method.

> **NOTE**
> For a package method to be accessed externally, it has to begin with an uppercase

The `main.go` file under package main, imports our custom package and calls the `Init` method in our `routes` package.

## Usage

### Without Docker

* Clone this repo: `git clone https://github.com/ichtrojan/go-package-example.git`
* Change directory: `cd go-package-example`
* Run the application: `go run main.go`

You will be able to access the application from `localhost:9990`

### With Docker

This is the good part, just run:

`docker-compose up`

You will be able to access the application from `localhost:9990`

viola, that all 😀
