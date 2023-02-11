/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewTfCommand() *cobra.Command {
	const StdIn = "-"
	var command = &cobra.Command{
		Use:   "tf <action>",
		Short: "Run argocd-tf-plugin",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("<action> argument required to acion command")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			if path == StdIn {
				
			} 
			return nil
		},
	}

	return command
}
