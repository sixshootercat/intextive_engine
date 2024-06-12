package utils

type Index map[string][]int

// Add adds the given documents to the index.
// It analyzes the text of each document and adds the document ID to the index for each token.
// If a document is already in the index for a particular token, it is skipped.
func (idx Index) Add(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// already in index
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

// intersection returns a new slice that contains the common elements between two input slices.
// The input slices `a` and `b`. The function iterates through both slices simultaneously and adds
// the common elements to the result slice. The resulting slice is also sorted in ascending order.
func intersection(a, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// Search searches for the given text in the index and returns a list of matching document IDs.
// If the text is not found in the index, it returns nil.
func (idx Index) Search(text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			// token not found
			return nil
		}
	}
	return r
}
