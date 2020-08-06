package util

import (
	"conflr/common"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/pkg/errors"
)

const (
	CONTENT_API_PATH = "rest/api/content"
)

func FetchArticlesStr(endpoint, expand, username, password string, from, limit int) (string, error) {

	// parse endpoint
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", errors.Wrap(err, "Invalid endpoint")
	}
	u.Path = path.Join(u.Path, CONTENT_API_PATH)

	// build URL
	q := u.Query()
	q.Set("type", common.CONFL_CONTENT_TYPE)
	q.Set("start", strconv.Itoa(from))
	q.Set("limit", strconv.Itoa(limit))
	q.Set("expand", expand)
	u.RawQuery = q.Encode()

	// build request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", errors.Wrap(err, "Failed to make request")
	}
	req.Header.Add("Accept", CONTENT_TYPE_JSON)
	req.Header.Add("Authorization", toBasicAuthHeader(username, password))

	// fetch articles (string)
	client := new(http.Client)
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return "", errors.Wrap(err, "failed to fetch data")
	} else if res.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("HTTP(%d)", res.StatusCode))
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read response data")
	}

	return string(bytes), nil
}

func toBasicAuthHeader(username, password string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return fmt.Sprintf("Basic %s", encoded)
}
