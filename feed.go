package atomus

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	atomXMLNS       = "http://www.w3.org/2005/Atom"
	atomContentType = "xhtml"
)

type Feed struct {
	XMLName      xml.Name `xml:"feed"` // <feed xmlns="http://www.w3.org/2005/Atom">
	XMLNamespace string   `xml:"xmlns,attr"`
	Title        string   `xml:"title"` // <title>Example Feed</title>
	//Subtitle     string   `xml:"subtitle,omitempty"`
	Link    *Link  `xml:"link,omitempty"` //<link href="http://example.org/" />
	Updated string `xml:"updated"`        // <updated>2003-12-13T18:30:02Z</updated>
	Entry   *Entry `xml:"entry"`
}

type Link struct {
	XMLName xml.Name `xml:"link"`
	Href    string   `xml:"href,attr"`
}

func NewFeed(title, link string) *Feed {
	return &Feed{
		XMLNamespace: atomXMLNS,
		Title:        title,
		Link: &Link{
			Href: link,
		},
	}
}

func (f *Feed) SetEntry(title, author, content string) {
	updated := time.Now()
	updatedRFC3339 := updated.Format(time.RFC3339)

	f.Updated = updatedRFC3339
	f.Entry = &Entry{
		Title:     fmt.Sprintf("%s - %s", title, updated.Format(time.RFC1123)),
		Id:        strconv.FormatInt(updated.Unix(), 10),
		Published: updatedRFC3339,
		Updated:   updatedRFC3339,
		Author: &EntryAuthor{
			Name: author,
		},
		Content: &EntryContent{
			Type:    atomContentType,
			Content: content,
		},
	}
}

func (f *Feed) Encode(w io.Writer) error {
	return xml.NewEncoder(w).Encode(f)
}
