/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "gogen",
	Short:             "A application generator for gRPC services in Go",
	Long:              "Gogen is a CLI tool generate the needed files to quickly create a gRPC service in Go.",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

// Execute runs the command.
func Execute() {
	_ = rootCmd.Execute()
}
