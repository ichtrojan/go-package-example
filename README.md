# Go Custom Package Example

## Introduction

This repo explain how to build and use your custom package in go.

The `routes/` directory contains a single file `api.go` under the package name `routes`. In that file `mux` router is imported; serving a single route to display `Hello World` declared inside of the `Init` method.

> **NOTE**
> For a package method to be accessed externally, it has to begin with an uppercase

The `main.go` file under package main, imports our custom package and calls the `Init` method in our `routes` package.

## Usage

### Without Docker

* Clone this repo
* Run `go get -u github.com/gorilla/mux`
* Run `go get -u github.com/ichtrojan/go-package-example/routes`
* Run `go run main.go`

You should be able to access the application from `localhost:9990`

### With Docker

This is the good part, just run:

`docker-compose up`

viola, that all 😀
