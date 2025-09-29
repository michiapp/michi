package handler

import (
	"net/http"

	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/gin-gonic/gin"
)

func (h *Handler) handleShortcut(ctx *gin.Context, action *parser.QueryAction) {
	result := action.Result

	if result == nil || len(result.Matches) == 0 {
		respondWithError(
			ctx,
			http.StatusInternalServerError,
			"handleShortcut: Received malformed parser result for query '%s'. Result: %+v.",
			"Failed to process your shortcut request. Please try again.",
			nil,
			ctx.Query(h.QueryParam), result,
		)
		return
	}

	alias := result.Matches[0]

	shortcut, err := h.services.GetShortcutService().GetFromAlias(ctx, alias)

	if err != nil {
		respondWithError(
			ctx,
			http.StatusInternalServerError,
			"handleShortcut: Error retrieving shortcut for alias '%s': %v",
			"Failed to retrieve shortcut. Please try again later.",
			err,
			alias,
		)
		return
	}

	ctx.Redirect(http.StatusFound, shortcut.Url)
}
