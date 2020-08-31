package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version info of server.",
	Long:  `Print the version info of server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("realifs version 1.0.0!")
	},
}
