package holdings

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdHoldings() *cobra.Command {
	holdingsCmd := &cobra.Command{
		Use:   "holdings",
		Short: "A brief description of your command",
		Long:  `A longer description`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("holdings called")
		},
	}
	return holdingsCmd
}
