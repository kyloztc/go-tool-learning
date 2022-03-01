package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of app demo",
	Long:  `All software has versions. This is version of app demo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("App demo version v1.0 -- HEAD")
	},
}
