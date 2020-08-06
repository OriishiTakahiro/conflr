package command

import (
	"conflr/common"
	"conflr/model"
	"conflr/util"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

var (
	importCmd = &cobra.Command{
		Use:   "import",
		Short: "import data from confl",
		RunE:  imporAllArticlestConflToSolr,
	}

	conflEndpointParam string
	solrEndpointParam  string
	usernameParam      string
	passwordParam      string
	fromParam          int
	limit              int
)

func init() {
	importCmd.PersistentFlags().StringVarP(
		&conflEndpointParam,
		"confl",
		"c",
		common.DEFAULT_CONFL_ENDPOINT,
		"confl=https://confl.example.com",
	)
	importCmd.PersistentFlags().StringVarP(
		&solrEndpointParam,
		"solr",
		"s",
		common.DEFAULT_SOLR_ENDPOINT,
		"endpoint=https://localhost:8983",
	)
	importCmd.PersistentFlags().StringVarP(
		&usernameParam,
		"user",
		"u",
		"username",
		"user=username",
	)
	importCmd.PersistentFlags().StringVarP(
		&passwordParam,
		"pass",
		"p",
		"password",
		"pass=password",
	)
	importCmd.PersistentFlags().IntVarP(
		&fromParam,
		"from",
		"f",
		common.DEFAULT_FROM,
		"from=0",
	)
	importCmd.PersistentFlags().IntVarP(
		&limit,
		"limit",
		"l",
		common.DEFAULT_LIMIT,
		"limit=100",
	)
}

func imporAllArticlestConflToSolr(cmd *cobra.Command, args []string) error {
	currentFrom := fromParam
	var err error
	for {
		log.Printf("Start importing articles from %d\n", currentFrom)
		currentFrom, err = importConflToSolr(
			conflEndpointParam,
			common.EXPAND,
			usernameParam,
			passwordParam,
			currentFrom,
			limit,
		)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to import articles from(%d) limit(%d)", currentFrom, limit)
			errors.Wrap(err, errMsg)
		}
		if currentFrom < 0 {
			log.Println("All articles are imported!")
			return nil
		}
		log.Printf("Impoted articles until %d\n", currentFrom-1)
	}
	return nil
}

func importConflToSolr(conflEndpoint, expand, username, password string, from, limit int) (int, error) {

	// fetch articles from confl
	artiConfList, err := fetchArticlesConfl(
		conflEndpoint,
		common.EXPAND,
		username,
		password,
		from,
		limit,
	)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to fetch confl's articles")
	}

	// convert confl's articles to solr's articles
	artiSolrList := make([]model.Article, len(artiConfList.Results))
	for i, v := range artiConfList.Results {
		artiSolrList[i] = model.NewArticle(v, conflEndpoint)
	}
	if len(artiSolrList) < limit {
		return -1, nil
	}

	// remove HTML tags
	for i, v := range artiSolrList {
		plainView, err := util.HTML2PlainText(v.View)
		if err != nil {
			return -1, errors.Wrap(err, "Failed to remove HTML tags")
		}
		artiSolrList[i].View = plainView
	}

	// indexing articles
	err = indexingArticle(common.DEFAULT_SOLR_ENDPOINT, common.DEFAULT_SOLR_CORE, true, artiSolrList)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to indexing")
	}

	// display last ID
	maxID := from
	for _, v := range artiSolrList {
		id, err := strconv.Atoi(v.ID)
		if err != nil {
			return -1, errors.Wrap(err, "Failed to convert article's ID to int")
		}
		if maxID < id {
			maxID = id
		}
	}

	// return next from paramter
	return artiConfList.NextFrom()
}

func fetchArticlesConfl(endpoint, expand, username, password string, from, limit int) (model.ArticleConflList, error) {

	// fetch article string from confl
	artiListStr, err := util.FetchArticlesStr(endpoint, expand, username, password, from, limit)
	if err != nil {
		return model.ArticleConflList{}, errors.Wrap(err, "Failed to featch Article string")
	}

	// convert json ArticleConflList
	var artiList model.ArticleConflList
	if err := json.Unmarshal([]byte(artiListStr), &artiList); err != nil {
		return model.ArticleConflList{}, errors.Wrap(err, "Failed to parse JSON string")
	}

	return artiList, nil
}

func indexingArticle(endpoint, core string, commit bool, articles []model.Article) error {

	bytes, err := json.Marshal(articles)
	if err != nil {
		return errors.Wrap(err, "Failed to convert to json")
	}

	return util.IndexingSolr(endpoint, core, commit, string(bytes))
}
