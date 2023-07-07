package shrls

import (
	"context"
	"fmt"

	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/config"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/service"
	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ShrlsCmd.AddCommand(shrlsCreateCmd)
}

var ShrlsCmd = &cobra.Command{
	Use:   "shrl",
	Short: "Manage shortened urls via command line",
	Long:  `Perform CRUD operations on shortened urls within Shrls`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()
	},
}

var shrlsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a new short url to Shrls",
	Long:  `Create a new short url and add it to shrls.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		defer ctx.Done()

		config := config.New()
		s := service.New(config)
		client := s.NewClient()
		shrl, err := client.CreateShrl(ctx, &pb.ShortURL{
			Stub: "hello",
			Content: &pb.ExpandedURL{
				Content: &pb.ExpandedURL_Url{
					Url: &pb.Redirect{
						Url: "https://www.google.com/2",
					},
				},
			},
			Tags: []string{"wild", "twos"},
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Info("Url shortened", fmt.Sprintf("Shrl: %#v", shrl))
	},
}
