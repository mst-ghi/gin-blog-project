package main

import (
	"blog/core/cmd"
)

// @securityDefinitions.apikey Bearer
// @in 						   header
// @name 					   Authorization
func main() {
	cmd.Execute()
}
