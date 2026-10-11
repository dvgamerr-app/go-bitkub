package cmd

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

var stdJson = jsoniter.ConfigCompatibleWithStandardLibrary

func pageFlags(cmd *cobra.Command) (page int, limit int) {
	page, _ = cmd.Flags().GetInt("page")
	limit, _ = cmd.Flags().GetInt("limit")
	return max(page, 0), max(limit, 0)
}
