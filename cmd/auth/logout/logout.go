package logout

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdLogout() *cobra.Command {
	logoutCmd := &cobra.Command{
		Use:   "logout",
		Short: "A brief description of your command",
		Long: `A longer description`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("logout called")
		},
	}
	return logoutCmd
}
