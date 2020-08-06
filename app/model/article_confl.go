package model

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

type (
	Space struct {
		Name string `json:"name"`
	}

	History struct {
		CreatedBy struct {
			UserName    string `json:"username"`
			DisplayName string `json:"displayName"`
		} `json:"createdBy"`
	}

	Body struct {
		View struct {
			Value string `json:"value"`
		} `json:"view"`
	}

	ArticleConfl struct {
		ID      string  `json:"id"`
		Type    string  `json:"type"`
		Title   string  `json:"title"`
		Space   Space   `json:"space"`
		History History `json:"history"`
		Body    Body    `json:"body"`
		Links   struct {
			WebUI string `json:"webui"` // webUI path
		} `json:"_links"`
	}

	ArticleConflList struct {
		Results []ArticleConfl `json:"results"`
		Links   struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"_links"`
	}
)

func (r ArticleConflList) NextFrom() (int, error) {
	u, err := url.Parse(r.Links.Next)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse URL %s", r.Links.Next)
		return -1, errors.Wrap(err, msg)
	}

	nextFromStr := u.Query().Get("start")
	nextFrom, err := strconv.Atoi(nextFromStr)
	if err != nil {
		msg := fmt.Sprintf("Failed to convert to int (%s)", nextFromStr)
		return -1, errors.Wrap(err, msg)
	}

	return nextFrom, nil
}
