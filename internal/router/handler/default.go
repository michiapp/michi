package handler

import (
	"net/http"

	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/gin-gonic/gin"
)

func (h *Handler) handleDefaultSearch(ctx *gin.Context, action *parser.QueryAction) {
	if action.RawQuery == "" {
		respondWithError(
			ctx,
			http.StatusBadRequest,
			"handleDefaultSearch: RawQuery is empty for default search type. Action: %+v",
			"We couldn't find a valid query for your search.",
			nil,
			action,
		)
		return
	}

	p, url, err := h.services.GetProvidersService().ResolveWithFallback(ctx, action.RawQuery)

	if err != nil || url == nil {
		providerTag := "N/A"
		respondWithError(
			ctx,
			http.StatusInternalServerError,
			"handleDefaultSearch: Failed to resolve redirect for query '%s' with provider '%s': %v",
			"Could not determine search destination. Please try again later.",
			err,
			action.RawQuery, providerTag,
		)
		return
	}

	h.completeSearchRequest(ctx, *url, action.Result, p)
}
