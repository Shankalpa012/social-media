package cmd

import (
	"fmt"

	"github.com/shankalpa12/cmd/bootstrap"
	"github.com/spf13/cobra"
)

var (
	webCmd = &cobra.Command{
		Use:   "practice:server",
		Short: "starting the bootstrap of web http server",
		Long:  "From this point http server is called and routes and end points are called",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Server Starting")
			bootstrap.WebApplication()
		},
	}
)

func init() {
	rootCmd.AddCommand(webCmd)
}
