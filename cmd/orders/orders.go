package orders

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdOrders() *cobra.Command {
	ordersCmd := &cobra.Command{
		Use:   "orders",
		Short: "A brief description of your command",
		Long:  `A longer description`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("orders called")
		},
	}
	return ordersCmd
}
