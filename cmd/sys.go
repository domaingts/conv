/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var (
	start   bool
	stop    bool
	restart bool
	reload  bool
)

// sysCmd represents the sys command
var sysCmd = &cobra.Command{
	Use:   "sys",
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cm := sysRun(ctx, args[0])
		fmt.Println(cm.String())
		err := cm.Run()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sysCmd)

	sysCmd.Flags().BoolVarP(&start, "start", "a", false, "Start")
	sysCmd.Flags().BoolVarP(&stop, "stop", "o", false, "Stop")
	sysCmd.Flags().BoolVarP(&restart, "restart", "s", false, "Restart")
	sysCmd.Flags().BoolVarP(&reload, "reload", "r", false, "Reload")
}

func sysRun(ctx context.Context, com string) *exec.Cmd {
	var verb string
	switch {
	case start:
		verb = "start"
	case stop:
		verb = "stop"
	case restart:
		verb = "restart"
	case reload:
		verb = "reload"
	default:
		verb = "status"
	}
	return exec.CommandContext(ctx, "systemctl", verb, com)
}
