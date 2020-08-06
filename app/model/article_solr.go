package model

import (
	"conflr/common"
	"strings"
)

type Article struct {
	ID                   string   `json:"id"`
	Type                 string   `json:"type"`
	Title                string   `json:"title"`
	SpaceName            string   `json:"space_name"`
	CreatedByUserName    string   `json:"createdBy_username"`
	CreatedByDisplayName string   `json:"createdBy_displayName"`
	View                 string   `json:"view"`
	URL                  string   `json:url`
	Labels               []string `json:"labels"`
}

func NewArticle(a ArticleConfl, endpoint string) Article {
	art := Article{
		ID:                   a.ID,
		Type:                 a.Type,
		Title:                a.Title,
		SpaceName:            a.Space.Name,
		CreatedByUserName:    a.History.CreatedBy.UserName,
		CreatedByDisplayName: a.History.CreatedBy.DisplayName,
		View:                 a.Body.View.Value,
		URL:                  endpoint + a.Links.WebUI,
		Labels:               []string{},
	}

	for k, terms := range common.LABEL_TERMS {
		for _, t := range terms {
			if strings.Contains(art.Title, t) {
				art.Labels = append(art.Labels, k)
				break
			}
		}
	}

	return art
}
