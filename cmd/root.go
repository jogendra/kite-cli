package cmd

import (
	"fmt"
	"kite/cmd/auth"
	"kite/cmd/gtt"
	"kite/cmd/holdings"
	"kite/cmd/orders"
	"kite/cmd/positions"
	"kite/cmd/profile"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kite",
	Short: "A brief description of your application",
	Long:  `A longer description`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Add Subcommands
	addSubcommands()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kite.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".kite" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".kite")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func addSubcommands() {
	authCmd := auth.NewCmdAuth()
	rootCmd.AddCommand(authCmd)
	gttCmd := gtt.NewCmdGTT()
	rootCmd.AddCommand(gttCmd)
	holdingsCmd := holdings.NewCmdHoldings()
	rootCmd.AddCommand(holdingsCmd)
	positionsCmd := positions.NewCmdPositions()
	rootCmd.AddCommand(positionsCmd)
	ordersCmd := orders.NewCmdOrders()
	rootCmd.AddCommand(ordersCmd)
	profileCmd := profile.NewCmdProfile()
	rootCmd.AddCommand(profileCmd)
}
