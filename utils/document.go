package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// LoadDocuments loads documents from the specified path.
// It reads an XML file, decompresses it if necessary, and decodes the XML into a slice of documents.
// Each document is assigned a unique ID starting from 0.
// The function returns the loaded documents and any error encountered during the process.
func LoadXMLDocuments(path string) (docs []document, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	xmlDecoder := xml.NewDecoder(gz)
	dump := struct {
		Documents []document `xml:"doc"`
	}{}

	if err := xmlDecoder.Decode(&dump); err != nil {
		return nil, err
	}

	docs = dump.Documents

	for i := range docs {
		docs[i].ID = i
	}

	return docs, nil
}
