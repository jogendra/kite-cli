package positions

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdPositions() *cobra.Command {
	positionsCmd := &cobra.Command{
		Use:   "positions",
		Short: "A brief description of your command",
		Long:  `A longer description`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("positions called")
		},
	}
	return positionsCmd
}
