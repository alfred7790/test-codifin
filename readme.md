# Codifin-challenge
This project has been deployed at [API server](https://codifin-challenge.mi-escaparate.com/v1/swagger/index.html) and will be available until March 1st of 2024.

```shell
Autor: I.S.C. Edgar Alfred Rodriguez Robles
E-mail: alfred.7790@gmail.com
```
# Overview
This project is an example of API REST.

## Dependencies

- [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) v1.9.1
- [github.com/go-gormigrate/gormigrate/v2](https://github.com/go-gormigrate/gormigrate/v2) v2.1.1
- [github.com/jinzhu/configor](https://github.com/jinzhu/configor) v1.2.1
- [github.com/swaggo/files](https://github.com/swaggo/files) v1.0.1
- [github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger) v1.6.0
- [github.com/swaggo/swag](https://github.com/swaggo/swag) v1.16.3
- [gorm.io/driver/postgres](https://github.com/gorm.io/driver/postgres) v1.5.4
- [gorm.io/gorm](https://github.com/gorm.io/gorm) v1.25.5

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
[GIN-debug] Listening and serving HTTP on :1315
```
7. Go to [swagger docs](http:localhost:1315/v1/swagger/index.html) and have fun.

# Testing
- Using `Makefile`.
```shell
$ make test
```
- Or try this:
```shell
$ go test ./...
```

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
[GIN-debug] Listening and serving HTTP on :1315
```
9. Go to [swagger docs](http:localhost:1315/v1/swagger/index.html) and have fun.

# Custom Config
> If you need to change the default values of the configuration.
1. Open the project.
```shell
$ cd test-codifin
```
2. Edit the configuration file `config/yaml/config_develop.yml`.

> WARNING! Make sure that if you edit the values about the `DB service`, also you should modify the `docker-compose.yml` file.

3. Restart the service `using Makefile` or `running Binary`.
4. Go to [swagger docs](http:localhost:1315/v1/swagger/index.html) and have fun.

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