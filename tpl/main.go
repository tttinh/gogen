package tpl

func MainTemplate() []byte {
	return []byte(`
package main

import "{{ .PackageName }}/cmd"

func main() {
	cmd.Execute()
}
`)
}

func RootTemplate() []byte {
	return []byte(`
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "{{ .ApplicationName }}",
	Short: "A brief description of your application",
	Long: ` + "`" + `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

XYZ is a service that do some magical things.` + "`" + `,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
`)
}
