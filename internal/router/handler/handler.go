package handler

import (
	"fmt"
	"net/http"

	"github.com/OrbitalJin/michi/internal"
	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/sqlc"
	"github.com/OrbitalJin/michi/public"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	QueryParam  string
	queryParser parser.QueryParserIface
	services    *service.Services
	config      *internal.Config
}

func NewHandler(
	config *internal.Config,
	qp parser.QueryParserIface,
	services *service.Services,
	queryParam string,
) *Handler {

	return &Handler{
		config:      config,
		queryParser: qp,
		services:    services,
		QueryParam:  queryParam,
	}
}

func (h *Handler) Favicon(ctx *gin.Context) {
	data, err := public.Content.ReadFile("assets/favicon.svg")
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.Data(http.StatusOK, "image/svg+xml", data)
}

func (h *Handler) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func (h *Handler) Error(ctx *gin.Context) {
	message := ctx.Query("message")

	if message != "" {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": message})
	}

	ctx.HTML(http.StatusInternalServerError, "error.html", nil)
}

func (h *Handler) SessionOpened(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "session_success.html", nil)
}

func (h *Handler) completeSearchRequest(
	ctx *gin.Context,
	redirectURL string,
	result *parser.Result,
	provider sqlc.SearchProvider,
) {

	ctx.Redirect(http.StatusFound, redirectURL)

	fmt.Println("foo bar")
	fmt.Println(h.config.Service.History)
	if h.config.Service.History {
		go h.logSearchHistoryAsync(ctx, result, provider)
	}
}
