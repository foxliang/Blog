package main

import "go-mysql/route"

func main() {
	// start api server
	route.NewServer().Start()
}
