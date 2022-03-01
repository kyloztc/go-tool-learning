package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var EchoTimes int

var echoCmd = &cobra.Command{
	Use: "echo",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < EchoTimes; i++ {
			fmt.Println("hello world")
		}
	}}

func init() {
	echoCmd.Flags().IntVarP(&EchoTimes, "times", "t", 1, "times of echo")
	rootCmd.AddCommand(echoCmd)
}
