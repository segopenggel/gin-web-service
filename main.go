package main

import "gin-web-app/router"

func main() {
	router := router.InitRouter()
	router.Run(":80")
}
