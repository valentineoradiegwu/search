package query

import (
	"log"
	"search/match"
	"search/sources"
	"sync"
)

func Find(key string) {
	log.Printf("Searching for key %s", key)
	feeds, err := sources.GetFeeds()
	log.Printf("Found %d feeds", len(feeds))

	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(feeds))

	results := make(chan *match.Result)

	for _, feed := range feeds {
		matcher, err := match.GetMatcher(feed.Type)
		if err != nil {
			matcher, err = match.GetMatcher("default")
			if err != nil {
				log.Println(err)
				continue
			}
		}

		go func(matcher match.Matcher, feed *sources.Feed) {
			match.Match(matcher, feed, key, results)
			wg.Done()
		}(matcher, feed)

	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		log.Printf("Field: %s \nContent: %s", res.Field, res.Contents)
	}
}
