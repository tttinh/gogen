/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/tinhtt/gogen/fs"
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
		projectPath, err := generateProjectStructure()
		errorExit(err)

		getDependencies()

		log.Println("Your application is ready at ", projectPath)
	},
}

func errorExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateProjectStructure() (string, error) {
	log.Println("Generating project structure...")
	return fs.Generate()
}

func getDependencies() {
	log.Println("Getting project dependencies...")
	errorExit(goGet("github.com/spf13/cobra"))
	errorExit(goGet("github.com/spf13/viper"))
}

func goGet(mod string) error {
	log.Println("go get ", mod)
	return exec.Command("go", "get", mod).Run()
}

func init() {
	rootCmd.AddCommand(initCmd)
}
