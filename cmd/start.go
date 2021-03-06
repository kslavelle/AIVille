package main

import "github.com/kslavelle/AIVille/pkg/router"

func main() {
	apiRouter := router.CreateAPI()
	apiRouter.Run(":9001")
}
