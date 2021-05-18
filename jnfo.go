package jnfo

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Jnfo struct {
	Num        string
	Title      string
	Date       string
	Duration   string
	Director   string
	Studio     string
	Label      string
	Serie      string
	PicLink    string
	Categories []string
	Cast       []string
}

var minUnit = regexp.MustCompile(`[^0-9]`)

func New(url string) (*Jnfo, error) {
	doc, err := getDoc(url)
	if err != nil {
		return nil, err
	}

	nfo := &Jnfo{}

	nfo.PicLink, _ = doc.Find(".bigImage img").Attr("src")
	nfo.Title, _ = subStrAfterSpace(doc.Find(".container h3").Text())
	nfo.setNumDateDuration(doc)
	nfo.setMeta(doc)
	nfo.setCategories(doc)
	nfo.setCast(doc)

	return nfo, nil
}

func (nfo *Jnfo) setNumDateDuration(doc *goquery.Document) {
	doc.Find(".info p").Each(func(i int, s *goquery.Selection) {
		entry := s.Text()
		content, exists := subStrAfterSpace(entry)
		if !exists {
			return
		}

		if i == 0 {
			nfo.Num = content
			return
		}

		if i == 1 {
			nfo.Date = content
			return
		}

		if i == 2 {
			nfo.Duration = minUnit.ReplaceAllString(content, "")
			return
		}
	})
}

func (nfo *Jnfo) setMeta(doc *goquery.Document) {
	doc.Find(".info p a").Each(func(i int, s *goquery.Selection) {
		href := s.AttrOr("href", "")
		content := s.Text()

		if strings.Index(href, "director") > 0 {
			nfo.Director = content
			return
		}

		if strings.Index(href, "studio") > 0 {
			nfo.Studio = content
			return
		}

		if strings.Index(href, "label") > 0 {
			nfo.Label = content
			return
		}

		if strings.Index(href, "series") > 0 {
			nfo.Serie = content
			return
		}
	})
}

func (nfo *Jnfo) setCategories(doc *goquery.Document) {
	elements := doc.Find(".genre label a")
	cates := make([]string, 0, elements.Length())
	elements.Each(func(i int, s *goquery.Selection) {
		cates = append(cates, s.Text())
	})
	nfo.Categories = cates
}

func (nfo *Jnfo) setCast(doc *goquery.Document) {
	elements := doc.Find(".genre > a")
	cast := make([]string, 0, elements.Length())
	elements.Each(func(i int, s *goquery.Selection) {
		cast = append(cast, s.Text())
	})
	nfo.Cast = cast
}
