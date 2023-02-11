/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewTfCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "tf",
		Short: "Run argocd-tf-plugin",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "argocd-tf-plugin %s ", version)
		},
	}

	return command
}
