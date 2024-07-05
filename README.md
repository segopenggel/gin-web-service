gin-web-app
======
Many thanks to this repo : https://github.com/yanyumiao/gin-demo 
##### About
A demo of developing web services using gin framework
* MVC layering
* Uniform output format
* Support MySQL Redis
##### Start
1 DB Preparation
* demo using MySQL database
* Create test database first and run person.sql to create person table
* Note that redis has set a password. See the setting in redis.go file is "123456"

2 Go Environment Setup

The project has been updated and now uses go module for project management
So before starting the project, you need to set up support module

* Necessary condition go version >= v1.11
* Set up go support module
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```
Note:

Because go module is used to manage the project, the project location should no longer be placed in $GOPATH/src

3 Run
Because go module is used, run it directly in the directory where main.go is located
```
go run main.go
```
Visit 127.0.0.1:10086 and see hello output Success!