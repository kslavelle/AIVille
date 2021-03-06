package main

import "github.com/kslavelle/AIVille/pkg/router"

func main() {
	apiRouter, connection := router.CreateAPI()
	defer connection.Close()

	apiRouter.Run(":9001")
}
