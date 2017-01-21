package topic

import (
	"github.com/ovh/tat"
	"github.com/ovh/tat/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdTopicTruncateTags = &cobra.Command{
	Use:   "truncatetags",
	Short: "Truncate Tags on this topic, only for tat admin and administrators on topic : tatcli topic truncatetags <topic>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			out, err := internal.Client().TopicTruncateTags(tat.TopicNameJSON{Topic: args[0]})
			internal.Check(err)
			if internal.Verbose {
				internal.Print(out)
			}
		} else {
			internal.Exit("Invalid argument: tatcli topic truncatetags --help\n")
		}
	},
}
