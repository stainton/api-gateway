/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/stainton/api-gateway/cmd"

func main() {
	gatewayCmd := cmd.NewCmdGateway()
	gatewayCmd.Execute()
}
