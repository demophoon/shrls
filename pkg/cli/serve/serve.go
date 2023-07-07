package serve

import (
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/config"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/service"

	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the Shrls server.",
	Long:  `Start the Shrls server using the settings given.`,
	Run:   shrls_serve,
}

func shrls_serve(cmd *cobra.Command, args []string) {
	config := config.New()
	s := service.New(config)
	s.Run()
}
