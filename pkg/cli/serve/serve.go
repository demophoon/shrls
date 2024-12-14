package serve

import (
	"os"
	"path"

	"github.com/demophoon/shrls/pkg/config"
	"github.com/demophoon/shrls/pkg/service"
	"github.com/demophoon/shrls/pkg/version"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the Shrls server.",
	Long:  `Start the Shrls server using the settings given.`,
	Run:   shrls_serve,
}

func shrls_serve(cmd *cobra.Command, args []string) {
	if viper.GetBool("demo") {
		// Create temp directories
		shrls_demo, err := os.MkdirTemp("", "shrls-demo")
		if err != nil {
			log.Fatalf("Unable to create demo dir: %s", err)
		}
		defer os.RemoveAll(shrls_demo)
		db_path, err := os.CreateTemp(shrls_demo, "shrls.db")
		if err != nil {
			log.Fatalf("Unable to create boltdb: %s", err)
		}
		// Initialize Configuration
		viper.Set("state.bolt.path", db_path.Name())
		viper.Set("uploads", path.Join(shrls_demo, "uploads"))
		viper.Set("default_redirect", "/admin/")

		log.Infof("Demo Mode initialized in %s", shrls_demo)

		setupDemo()
	}

	log.Info("SHRLS ", version.Version)
	config := config.New()
	s := service.New(config)
	s.Run()
}
