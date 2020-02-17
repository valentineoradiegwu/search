package match

import (
	"search/sources"
)

type DefaultMatcher struct {
}

func (r *DefaultMatcher) Search(feed *sources.Feed, key string) ([]*Result, error) {
	return nil, nil
}
