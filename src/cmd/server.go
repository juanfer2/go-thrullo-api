package cmd

import (
	"github.com/juanfer2/go-thrullo-api.git/src/servers"
	"github.com/spf13/cobra"
)

var ServerCmd = cobra.Command{
	Use:  "server",
	Long: "Migrate database strucutures. This will create new tables and add missing columns and indexes.",
	Run:  startServer,
}

func startServer(cmd *cobra.Command, args []string) {
	servers.StartServer()
}
