package match

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"search/sources"
	"strings"
)

type RSSMatcher struct {
}

type rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel channel  `xml:"channel"`
}

type item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
}

type channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Items       []item   `xml:"item"`
}

func (r *RSSMatcher) Search(feed *sources.Feed, key string) ([]*Result, error) {
	if feed.Link == "" {
		return nil, fmt.Errorf("The URL for the link is empty")
	}
	log.Printf("Downloading feed %s", feed.Link)
	resp, err := http.Get(feed.Link)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Expecting a 200 response code  but received %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssXML rss
	err = xml.Unmarshal(body, &rssXML)

	if err != nil {
		return nil, err
	}

	var matches []*Result
	for _, item := range rssXML.Channel.Items {
		if strings.Contains(item.Title, key) {
			result := &Result{"Title", item.Title}
			matches = append(matches, result)
		}
		if strings.Contains(item.Description, key) {
			result := &Result{"Description", item.Description}
			matches = append(matches, result)
		}
	}

	log.Printf("Found %d matches for feed %s", len(matches), feed.Link)
	return matches, nil
}
