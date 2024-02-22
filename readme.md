# Codifin-challenge
```shell
Autor: I.S.C. Edgar Alfred Rodriguez Robles
E-mail: alfred.7790@gmail.com
```
# Table of Contents
[_TOC_]

# Overview
This project is an example of API REST.

# Requirements
To run this project it is necessary to have installed:
- docker:latest
- docker-compose:latest
- go 1.20
- swagger:latest

# GOPATH is exported?
Make sure that yout GOPATH is exported.
```shell
$ echo $GOPATH
```
> To create the API documentation, swagger will be installed in your go directory ($GOPATH).

# Quick start
> IMPORTANT! If you have a linux distribution, you will be able to perform this procedure, otherwise I suggest you try another way to run the service.
1. Clone the repo:
```shell
$ git clone git@github.com:alfred7790/test-codifin.git
```
2. Open the project
```shell
$ cd test-codifin
```
5. Build and run the service:
```shell
$ make deploy-local
```
6. If everything is ok, you should see something like this:
```shell
[GIN-debug] Listening and serving HTTP on :1313
```
7. Go to [swagger docs](http:localhost:1313/v1/swagger/index.html) and have fun.

# Manual start - Running Binary (optional)
1. Clone the repo:
```shell
$ git clone git@github.com:alfred7790/test-codifin.git
```
2. Open the project
```shell
$ cd test-codifin
```
3. Running the DB service:
```shell
$ docker-compose up -d db_products
```
4. Get dependencies:
```shell
$ go mod tidy && go get -u github.com/swaggo/swag/cmd/swag
```
5. Build the service:
```shell
$ go build -o build/bin/go_codifin cmd/codifin/main.go
```
6. Build swagger docs
```shell
swag init -g ./cmd/codifin/main.go
```
7. Running the service:
```shell
$ ./build/bin/go_codifin
```
8. If everything is ok, you should see something like this:
```shell
[GIN-debug] Listening and serving HTTP on :1313
```
9. Go to [swagger docs](http:localhost:1313/v1/swagger/index.html) and have fun.

# Custom Config
> If you need to change the default values of the configuration.
1. Open the project.
```shell
$ cd test-codifin
```
2. Edit the configuration file `config/yaml/config_develop.yml`.

> WARNING! Make sure that if you edit the values about the `DB service`, also you should modify the `docker-compose.yml` file.

3. Restart the service `using Makefile` or `running Binary`.
4. Go to [swagger docs](http:localhost:1313/v1/swagger/index.html) and have fun.

# Dependencies
> Make sure that your GOPATH is exported.

If you get an error with swagger or with another package, try this:
```shell
$ make dep
```
Or
```shell
$ go mod tidy && go get -u github.com/swaggo/swag/cmd/swag
```