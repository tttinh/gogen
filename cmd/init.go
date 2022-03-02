/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"initialize", "initialise", "create"},
	Short:   "Initialize a gRPC service",
	Long: `Initialize (gogen init) will create a new application, with the
appropriate structure for a gRPC service (follow DDD, Clean Architect and CQRS concepts).

Gogen init must be run inside of a go module (please run "go mod init <MODNAME>" first)
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
