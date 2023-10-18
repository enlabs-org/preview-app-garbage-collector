package cr

import (
	"enlabs/preview-app-garbage-collector/cmd/root"
	"enlabs/preview-app-garbage-collector/helpers/github"
	"github.com/spf13/cobra"
	"fmt"
)


var Cmd = &cobra.Command{
	Use:     "container-registry",
	Short:   "Clean container registry",
	Aliases: []string{"cr"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {

		pr,_ := github.GetPullRequests()
		fmt.Println(pr)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}