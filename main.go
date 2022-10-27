package main

import (
	"assignment-3/routers"
)

func main() {
	routers.StartServer().Run(":3000")
}
