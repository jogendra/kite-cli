package gtt

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdGTT() *cobra.Command {
	gttCmd := &cobra.Command{
		Use:   "gtt",
		Short: "A brief description of your command",
		Long:  `A longer description`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("gtt called")
		},
	}
	return gttCmd
}
