package main

import (
	"sports_api/router"
)

func main() {
	r := router.SetupRouter() // Initialize the main router
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
