package root

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "pa-garbage-collect",
	Short: "pa-garbage-collect is a CLI tool for cleaning up old preview apps",
}

func init() {}
