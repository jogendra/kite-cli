package login

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdLogin() *cobra.Command {
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "A brief description of your command",
		Long: `A longer description`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("login called")
		},
	}
	return loginCmd
}
