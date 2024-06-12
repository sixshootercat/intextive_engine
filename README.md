# Full Text Search Engine

A text search engine written in Go. Uses the [Wikipedia XML dump](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz) as the dataset source by default. Uses a simple inverted index to search for text in the dataset.

## High level overview

The entire dataset text is pre-processed and an inverted index is created after loading the documents in memory.

Example:

```json
documents = {
  1: "<abstract>hello world something</abstract>",
  2: "<abstract>hola mundo</abstract>",
  3: "<abstract>ciao mondo</abstract>",
  4: "<abstract>bye world</abstract>",
  5: "<abstract>lorem something else world</abstract>"
}

index = {
    "hello": [1]
    "world": [1 4 5]
    "hola": [2]
    "ciao": [3]
    "bye": [4]
    "lorem": [5]
    "something": [1 5]
    "else": [5]
}
```

## Why do we need a text search engine?

- to find the information we need quickly in a large amount of text and documents.
- they are optimized for speed and efficiency.
- they are designed to handle large amounts of data and provide fast search results. +10M documents can be searched sub second.

## How does it work?

The process of indexing a text document is as follows:

1. Tokenization: Break the text into individual words/tokens and remove any punctuation.
   Example:

```bash
tokenize("Hello world, how are you today?")
tokens = ["Hello", "world", "how", "are", "you", "today"]
```

2. Filtering:

   1. lowercase: make the search case-insensitive
   2. drop common words: stop words like "the", "and", "a", "an", "in", etc.
   3. stemming: reduce words to their base form (e.g. "running" to "run")

3. Indexing: Create an inverted index where each token is mapped to a list of document IDs that contain that token.
4. Search: Given a query, find the document IDs that contain all the tokens in the query.

## Running it

The program comes pre-configured with a gzipped sample XML dump of the Wikipedia documents and a sample query to search for. You can also provide the xml document path and search query as command flag values as shown below.

```bash
go run main.go -p "path/to/xml-dump.xml" -q "checkmate in two moves"
```
