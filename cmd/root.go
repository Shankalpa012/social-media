package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:              "Shankalpa Test API",
		Short:            "Practice For the Golang API",
		Long:             "I am Building this project for practicing Golang and building web APIs with it.Keep Supporting.Peace Out",
		TraverseChildren: true,
		Hidden:           true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
)

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Println("Error while Executing the Commnd")
		//exit command woth non zero means error and zero means error
		os.Exit(1)
	}
}
