package main

import (
	"github.com/EMZemskova/server_go/internal"
)

func main() {

	router := internal.GetRouters()
	router.Run("localhost:8080")
}
