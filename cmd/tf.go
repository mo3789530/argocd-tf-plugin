/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	myhcl "argocd-tf-plugin/pkg/hcl"
	"argocd-tf-plugin/pkg/terraform"
	"fmt"
	"log"
	"regexp"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/spf13/cobra"
)

func NewTfCommand() *cobra.Command {
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
			fmt.Println(args)
			dir, _ := cmd.Flags().GetString("dir")
			relative, _ := cmd.Flags().GetString("relative")

			switch args[0] {
			case "plan":
				return runPlan(dir, relative)
			case "apply":
				return runApply(dir, relative)
			default:
				log.Printf("%s not found \n", args[0])
			}

			return nil
		},
	}

	command.Flags().StringP("dirs", "d", "", "working dir")
	command.Flags().StringP("relative", "r", "", "relative")

	return command
}

func runPlan(dir string, relative string) error {
	folderpath := relative + "/" + dir
	tf, exec, err := initTf(folderpath)

	if err != nil {
		log.Printf("error terraform init %s", err)
		return err
	}

	isdiffer, err := exec.Plan(tf)
	if err != nil {
		log.Printf("err plan %s", err)
		return err
	}
	if !isdiffer {
		log.Printf("no changes")
	} else {
		show, err := exec.Show(tf, false)
		if err != nil {
			log.Printf("err show %s", err)
			return err
		}
		log.Println(show)
	}

	return nil
}

func runApply(dir string, relative string) error {
	folderpath := relative + "/" + dir
	tf, exec, err := initTf(folderpath)

	if err != nil {
		log.Printf("error terraform init %s", err)
		return err
	}

	err = exec.Apply(tf)
	if err != nil {
		log.Printf("terraform apply error %s", err)
		return err
	}
	return nil
}

func initTf(folderpath string) (*tfexec.Terraform, terraform.Exec, error) {
	filepath := folderpath + "/terraform.tf"

	log.Print(filepath)

	strVersion, err := myhcl.GetVersions(filepath)
	if err != nil {
		log.Fatalf("err %s", err)
		return nil, nil, err
	}
	rex := regexp.MustCompile("[0-9.]+")
	version := rex.FindString(strVersion)

	log.Printf("%s using terrafrom version: %s", filepath, version)

	exec := terraform.NewExec(version, true)
	tf, err := exec.Init(folderpath)
	if err != nil {
		log.Printf("error init %s", err)
		return nil, nil, err
	}
	return tf, exec, nil
}
