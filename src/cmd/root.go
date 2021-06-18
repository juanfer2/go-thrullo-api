package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile = ""

func init() {
	rootCmd.AddCommand(&MigrateCmd)
}

var rootCmd = cobra.Command{
	Use:   "gotrue",
	Short: "A generator for Cobra Kai based Applications",
	Long: `Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
}

func RootCommand() *cobra.Command {
	rootCmd.AddCommand(&MigrateCmd, &ServerCmd)
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	return &rootCmd
}
