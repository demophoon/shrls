package shrls

import (
	"context"
	"fmt"

	"github.com/demophoon/shrls/pkg/config"
	"github.com/demophoon/shrls/pkg/service"
	pb "github.com/demophoon/shrls/server/gen"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ShrlsCmd.AddCommand(shrlsCreateCmd)
	ShrlsCmd.AddCommand(shrlsListCmd)
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

var shrlsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List short urls",
	Long:  `List urls from the shrls server`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		defer ctx.Done()

		config := config.New()
		s := service.New(config)
		client := s.NewClient()

		shrls, _, err := client.ListShrls(ctx, nil, nil, nil)
		if err != nil {
			log.Fatal(err)
		}

		for _, shrl := range shrls {
			log.Info("Urls", fmt.Sprintf("%#v", shrl.Stub))
		}

	},
}
