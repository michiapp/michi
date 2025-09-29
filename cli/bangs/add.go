package bangs

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/sqlc"
	v2 "github.com/urfave/cli/v2"
)

func create(ctx context.Context, service *service.SPService) *v2.Command {

	siteNameFlag := &v2.StringFlag{
		Name:     "site-name",
		Aliases:  []string{"s"},
		Usage:    "The user-friendly name of the search provider (e.g., 'Google', 'DuckDuckGo Images')",
		Required: true,
	}

	tagFlag := &v2.StringFlag{
		Name:     "tag",
		Aliases:  []string{"t"},
		Usage:    "The short keyword or 'bang' for the provider (e.g., 'g', 'ddgi'). This is what you'd type to use it.",
		Required: true,
	}

	categoryFlag := &v2.StringFlag{
		Name:     "category",
		Aliases:  []string{"c"},
		Usage:    "The broad category of the provider (e.g., 'Web Search', 'Images', 'Dictionary')",
		Required: true,
	}

	subcategoryFlag := &v2.StringFlag{
		Name:     "subcategory",
		Aliases:  []string{"sc"},
		Usage:    "A more specific subcategory within its main category (e.g., 'General', 'Programming', 'News')",
		Required: true,
	}

	domainFlag := &v2.StringFlag{
		Name:     "domain",
		Aliases:  []string{"d"},
		Usage:    "The primary domain of the search provider (e.g., 'google.com', 'wikipedia.org')",
		Required: true,
	}

	urlFlag := &v2.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Usage: `The base URL for the search query. Use '%s' as a placeholder for the user's search terms.
               Example: For Google, use 'https://www.google.com/search?q=%s'`,
		Required: true,
	}

	return &v2.Command{
		Name:  "create",
		Usage: "Add a new search provider to the system.",
		Flags: []v2.Flag{
			siteNameFlag,
			tagFlag,
			categoryFlag,
			subcategoryFlag,
			domainFlag,
			urlFlag,
		},
		Action: func(c *v2.Context) error {
			siteName := c.String(siteNameFlag.Name)
			tag := c.String(tagFlag.Name)
			category := c.String(categoryFlag.Name)
			subcategory := c.String(subcategoryFlag.Name)
			domain := c.String(domainFlag.Name)
			url := c.String(urlFlag.Name)

			params := sqlc.InsertProviderParams{
				SiteName:    siteName,
				Tag:         tag,
				Category:    category,
				Subcategory: subcategory,
				Domain:      domain,
				Url:         url,
			}

			err := service.Insert(ctx, params)

			if err != nil {
				return fmt.Errorf("failed to add search provider: %w", err)
			}

			fmt.Printf("Successfully added search provider '%s' (tag: '%s')\n", params.SiteName, params.Tag)
			return nil
		},
	}
}
