package main

import (
	"flag"
	"log"
	"time"

	"github.com/sixshootercat/text_search_engine/utils"
)

func main() {
	var dumpPath, query string

	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "state variables", "search query")
	flag.Parse()

	start := time.Now()
	log.Println("Loading documents. This may take a while...")

	docs, err := utils.LoadXMLDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d documents in %s\n", len(docs), time.Since(start))

	start = time.Now()
	log.Println("Creating index. This may take a while...")
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %s\n\n", len(docs), time.Since(start))
	start = time.Now()

	matchedIDs := idx.Search(query)

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}

	log.Printf("Search found %d results in %s\n", len(matchedIDs), time.Since(start))
}
