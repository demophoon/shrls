package cli

import (
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the Shrls server.",
	Long:  `Start the Shrls server using the settings given.`,
	Run:   shrls_serve,
}

func shrls_serve(cmd *cobra.Command, args []string) {
	s := service.New()
	s.Run()
}
