package match

import (
	"fmt"
	"log"
	"search/sources"
	"sync"
)

type Matcher interface {
	Search(feed *sources.Feed, key string) ([]*Result, error)
}

var matchers map[string]Matcher
var match_rwmtx *sync.RWMutex

func init() {
	match_rwmtx = &sync.RWMutex{}
	matchers = make(map[string]Matcher)

	if err := RegisterMatcher("default", &DefaultMatcher{}); err != nil {
		log.Fatalf(err.Error())
	}

	if err := RegisterMatcher("rss", &RSSMatcher{}); err != nil {
		log.Fatalf(err.Error())
	}
}

func Match(matcher Matcher, feed *sources.Feed, key string, results chan<- *Result) {
	res, err := matcher.Search(feed, key)
	if err != nil {
		log.Println(err)
		return
	}

	for _, r := range res {
		results <- r
	}
}

func RegisterMatcher(name string, matcher Matcher) error {
	match_rwmtx.Lock()
	defer match_rwmtx.Unlock()

	if _, ok := matchers[name]; ok {
		return fmt.Errorf("A matching impl with name %s already exists", name)
	} else {
		matchers[name] = matcher
		return nil
	}
}

func GetMatcher(name string) (Matcher, error) {
	match_rwmtx.RLock()
	defer match_rwmtx.RUnlock()

	if impl, ok := matchers[name]; !ok {
		return impl, fmt.Errorf("A matching impl with name %s does not exists", name)
	} else {
		return impl, nil
	}
}
