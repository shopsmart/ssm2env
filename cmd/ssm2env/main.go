package main

import "github.com/shopsmart/ssm2env/cmd"

var version = "development"

func main() {
	cmd.Execute(version, nil)
}
