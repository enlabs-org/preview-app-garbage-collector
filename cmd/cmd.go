package cmd

import (
	"github.com/spf13/cobra"
	"enlabs/preview-app-garbage-collector/cmd/root"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}