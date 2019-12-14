package main

import "github.com/rchenhyy/demo-ginex/router"

func main() {
	r := router.SetupRouter()
	_ = r.Run()
}
