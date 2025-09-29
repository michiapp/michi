package handler

import (
	"net/http"

	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/gin-gonic/gin"
)

func (h *Handler) handleBang(ctx *gin.Context, action *parser.QueryAction) {
	result := action.Result
	service := h.services.GetProvidersService()

	if result == nil {
		respondWithError(
			ctx,
			http.StatusInternalServerError,
			"handleBang: Parser result is nil for query '%s'.",
			"An unexpected error occurred while processing your bang query.",
			nil,
			action.RawQuery,
		)
		return
	}

	best := service.Rank(ctx, result)

	provider, redirect, err := service.Resolve(
		result.Query,
		best,
	)

	if err != nil || redirect == nil {
		providerTag := "N/A"
		respondWithError(
			ctx,
			http.StatusInternalServerError,
			"handleBang: Failed to resolve redirect for query '%s' with provider '%s': %v",
			"Could not determine search destination. Please try again later.",
			err,
			result.Query, providerTag,
		)
		return
	}

	h.completeSearchRequest(ctx, *redirect, result, provider)
}
