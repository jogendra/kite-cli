package profile

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdProfile() *cobra.Command {
	profileCmd := &cobra.Command{
		Use:   "profile",
		Short: "A brief description of your command",
		Long:  `A longer description`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("positions called")
		},
	}
	return profileCmd
}
