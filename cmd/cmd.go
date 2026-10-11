package cmd

import (
	"github.com/dvgamerr-app/go-bitkub/crypto"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var stdJson = jsoniter.ConfigCompatibleWithStandardLibrary

func pageFlags(cmd *cobra.Command) (page int, limit int) {
	page, _ = cmd.Flags().GetInt("page")
	limit, _ = cmd.Flags().GetInt("limit")
	return max(page, 0), max(limit, 0)
}

func logPageLast(page int, last int) {
	log.Info().Int("page", page).Int("last", last).Msg("Pagination")
}

func runPaged[I any](cmd *cobra.Command, failure string, fetch func(page int, limit int) (int, int, []I, error), logItem func(I)) {
	page, limit := pageFlags(cmd)

	currentPage, lastPage, items, err := fetch(page, limit)
	if err != nil {
		log.Fatal().Err(err).Msg(failure)
	}

	logPageLast(currentPage, lastPage)

	for _, item := range items {
		logItem(item)
	}
}

func runHistory[I any](cmd *cobra.Command, failure string, fetch func(pagination crypto.Pagination, symbol string) (crypto.PaginatedResponse, []I, error), logItem func(I)) {
	page, limit := pageFlags(cmd)
	symbol, _ := cmd.Flags().GetString("symbol")

	paginated, items, err := fetch(crypto.Pagination{Page: page, Limit: limit}, symbol)
	if err != nil {
		log.Fatal().Err(err).Msg(failure)
	}

	logPagination(paginated)

	for _, item := range items {
		logItem(item)
	}
}
