package util

import (
	"strings"

	"github.com/pkg/errors"
	"jaytaylor.com/html2text"
)

func HTML2PlainText(html string) (string, error) {

	plainText, err := html2text.FromString(html, html2text.Options{PrettyTables: false, OmitLinks: true})
	if err != nil {
		return "", errors.Wrap(err, "cannot convert HTML to pain text")
	}

	plainText = trimMarkup(plainText)

	return plainText, nil
}

func trimMarkup(md string) string {
	var patterns = []string{"\\n", "*", "~", "#"}
	for _, v := range patterns {
		md = strings.Replace(md, v, "", -1)
	}

	return md
}
