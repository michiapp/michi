package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/sqlc"
)

func (h *Handler) logSearchHistoryAsync(ctx context.Context, result *parser.Result, provider sqlc.SearchProvider) {
	if result == nil {
		log.Printf(
			"logSearchHistoryAsync: Skipping history log due to missing result or provider. Result: %+v, Provider: %+v",
			result,
			provider,
		)
		return
	}

	params := sqlc.InsertHistoryEntryParams{
		Query:       result.Query,
		ProviderID:  provider.ID,
		ProviderTag: provider.Tag,
	}

	if err := h.services.GetHistoryService().Insert(ctx, params); err != nil {
		log.Printf(
			"failed to insert search history entry for query '%s': %v",
			params.Query,
			fmt.Errorf("insertion error: %w", err),
		)
	}
}
