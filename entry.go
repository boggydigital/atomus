package atomus

import (
	"encoding/xml"
)

type Entry struct {
	Title string `xml:"title"` // <title>Atom-Powered Robots Run Amok</title>
	// <link href="http://example.org/2003/12/13/atom03" />
	// <link rel="alternate" type="text/html" href="http://example.org/2003/12/13/atom03.html"/>
	// <link rel="edit" href="http://example.org/2003/12/13/atom03/edit"/>
	Id        string `xml:"id"`        // <id>urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a</id>
	Published string `xml:"published"` // <published>2003-11-09T17:23:02Z</published>
	Updated   string `xml:"updated"`   // <updated>2003-12-13T18:30:02Z</updated>
	// <summary>Some text.</summary>
	Author  *EntryAuthor  `xml:"author"`
	Content *EntryContent `xml:"content"`
}

type EntryContent struct {
	XMLName xml.Name `xml:"content"`
	Type    string   `xml:"type,attr"`
	Content string   `xml:",cdata"`
}

type EntryAuthor struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name"`            // <name>John Doe</name>
	Email   string   `xml:"email,omitempty"` // <email>johndoe@example.com</email>
}
