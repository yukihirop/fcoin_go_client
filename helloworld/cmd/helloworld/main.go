package main

import (
	"fcoin_go_client/helloworld"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

func main() {
	var prefix string
	var rootCmd = &cobra.Command{Use: "helloworld"}
	rootCmd.PersistentFlags().StringVar(&prefix, "prefix", "", "prefix for hello world")

	var cmdHelloWorld = &cobra.Command{
		Use:   "say",
		Short: "say hello world",
		Long:  "say hello world for test",
		Run: func(cmd *cobra.Command, args []string) {
			message := helloworld.CreateMessage(prefix)
			pp.Println(message)
		},
	}
	rootCmd.AddCommand(cmdHelloWorld)
	rootCmd.Execute()
}
