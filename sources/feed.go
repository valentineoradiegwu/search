package sources

import (
	"encoding/json"
	"io/ioutil"
)

type Feed struct {
	Site string `json:"site"`
	Link string `json:"link"`
	Type string `json:"type"`
}

func GetFeeds() ([]*Feed, error) {
	var feeds []*Feed
	b, err := ioutil.ReadFile("sources/data.json")

	if err != nil {
		return feeds, err
	}

	err = json.Unmarshal(b, &feeds)
	return feeds, err
}
