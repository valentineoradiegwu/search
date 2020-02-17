package main

import (
	"flag"
	"log"
	"os"
	"search/query"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	p_search_key := flag.String("search", "", "Specify the search key")
	flag.Parse()

	if *p_search_key == "" {
		log.Fatal("The 'search' command line argument is mandatory")
	}
	query.Find(*p_search_key)
}
