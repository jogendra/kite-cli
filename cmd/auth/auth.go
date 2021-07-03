package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"kite/cmd/auth/login"
	"kite/cmd/auth/logout"
	"kite/cmd/auth/refresh"
)

func NewCmdAuth() *cobra.Command {
	authCmd := &cobra.Command{
		Use:   "auth <command>",
		Short: "A brief description of your command",
		Long: `TO ADD LATER`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("auth called")
		},
	}
	// Add subcommands
	loginCmd := login.NewCmdLogin()
	authCmd.AddCommand(loginCmd)
	logoutCmd := logout.NewCmdLogout()
	authCmd.AddCommand(logoutCmd)
	refreshCmd := refresh.NewCmdRefresh()
	authCmd.AddCommand(refreshCmd)

	return authCmd
}
