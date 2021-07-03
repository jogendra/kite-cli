package refresh

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdRefresh() *cobra.Command {
	refreshCmd := &cobra.Command{
		Use:   "refresh",
		Short: "A brief description of your command",
		Long:  `A longer description`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("refresh called")
		},
	}
	return refreshCmd
}
