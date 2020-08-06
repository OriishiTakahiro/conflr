package util

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	PATH_TEMPLATE = "%s/solr/%s/update"
)

func IndexingSolr(endpoint, core string, commit bool, documentsStr string) error {

	// parse endpoint/path
	u, err := url.Parse(fmt.Sprintf(PATH_TEMPLATE, endpoint, core))
	if err != nil {
		return errors.Wrap(err, "Invalid URL")
	}

	// build URL
	q := u.Query()
	q.Set("commit", strconv.FormatBool(commit))
	u.RawQuery = q.Encode()

	// build request
	req, err := http.NewRequest("GET", u.String(), strings.NewReader(documentsStr))
	if err != nil {
		return errors.Wrap(err, "Failed to make request")
	}
	req.Header.Add("Content-type", CONTENT_TYPE_JSON)
	req.Header.Add("charset", CHAR_SET)

	log.Println(req.URL.String())
	log.Println(req.Header)
	log.Println(req.Body)

	// fetch articles (string)
	client := new(http.Client)
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return errors.Wrap(err, "failed to fetch data")
	} else if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("HTTP(%d)", res.StatusCode))
	}

	return nil
}
