package serve

import (
	"github.com/demophoon/shrls/pkg/config"
	"github.com/demophoon/shrls/pkg/service"
	"github.com/demophoon/shrls/pkg/version"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the Shrls server.",
	Long:  `Start the Shrls server using the settings given.`,
	Run:   shrls_serve,
}

func shrls_serve(cmd *cobra.Command, args []string) {
	log.Info("SHRLS ", version.Version)
	config := config.New()
	s := service.New(config)
	s.Run()
}
