/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	// "os"
	"os/exec"

	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic("empty args")
		}
		cm := exec.Command("journalctl", "-fu", args[0], "-o", "cat")
		cm.Stdout = os.Stdout
		cm.Stderr = os.Stderr
		fmt.Println(cm.String())
		err := cm.Start()
		if err != nil {
			panic(err)
		}
		err = cm.Wait()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
