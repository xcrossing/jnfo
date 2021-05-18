package jnfo

import (
	"errors"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getDoc(url string) (*goquery.Document, error) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	// Load the HTML document
	return goquery.NewDocumentFromReader(res.Body)
}

func subStrAfterSpace(str string) (string, bool) {
	strs := strings.SplitN(str, " ", 2)
	if len(strs) < 2 {
		return "", false
	}
	return strs[1], true
}
