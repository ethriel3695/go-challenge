package article

import "time"

type Quarter int

const (
	Unpublished Quarter = iota
	FirstQuarter
	SecondQuarter
	ThirdQuarter
	FourthQuarter
)

type Article struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Topics       []string  `json:"tags"`
	PublishMonth time.Time `json:"pubMonth"`
}

//
// Computes the fiscal quarter that the article was publish in
//
// 		FirstQuarter  - 01/01 - 03/31
// 		SecondQuarter - 04/01 - 06/31
// 		ThirdQuarter  - 07/01 - 09/30
// 		FourthQuarter - 10/01 - 12/31
//
// Returns an enum from 1-4
//
func (a *Article) Quarter() Quarter {
	// TODO - Implement

	return 0
}

//
// Compares the article to another article.
// The returned value is the number of topics that
// the two articles share, as well as a list of those related topics
//
func (a *Article) GetRelatedTopics(b *Article) (int, []string) {
	// TODO - Implement

	return 0, []string{}
}

//
// Returns an array of strings, where each string has the following format:
//
// 		"{Title} - Q{Quarter Number}"
//
// For example, given a list with an Article with the title, "How to get a personal loan"
// that was published in September, there should be a corresponding entry in the returned array
// with the value: "How to get a personal loan - Q3".
//
// If there are no entries for a quarter, there should be an entry in the
// returned array with the format:
//
// 		"No articles found for Q{Quarter Number}"
//
// For example, if there are no articles published in January, February, or March, there must
// be an entry in the returned array with the value: "No articles found for Q1".
//
func ListByTitleAndQuarter(articles []Article) []string {
	// TODO - Implement

	return []string{}
}

//
// Group the given articles by their topics
//
// The output is a map of articles with the key being the topic
// of all the contained articles
//
// {"topic": [{...}]}
//
func GroupByTopic(articles []Article) map[string][]Article {
	// TODO - Implement

	return map[string][]Article{}
}

//
// For a given article, return a list of articles that share at least
// one topic. The list is sorted with the most related article, to the
// least related article.
//
func FindRelatedArticles(needle Article, haystack []Article) []Article {
	// TODO - Implement

	return []Article{}
}