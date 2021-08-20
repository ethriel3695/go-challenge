package article

import (
	"github.com/google/go-cmp/cmp"
	"sort"
	"strings"
	"testing"
	"time"
)

/////////////////////////////////////////////////////
////	Unit Tests for Articles Package          ////
/////////////////////////////////////////////////////

// Problem 1
func TestQuarterGetter(t *testing.T) {
	testTable := map[string]struct {
		PublishMonth time.Time
		Expected     Quarter
	}{
		"can parse Q1":  {PublishMonth: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), Expected: FirstQuarter},
		"can parse Q2":  {PublishMonth: time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC), Expected: SecondQuarter},
		"can parse Q3":  {PublishMonth: time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC), Expected: ThirdQuarter},
		"can parse Q4":  {PublishMonth: time.Date(2019, 11, 1, 0, 0, 0, 0, time.UTC), Expected: FourthQuarter},
		"can parse nil": {PublishMonth: time.Time{}, Expected: Unpublished},
	}

	for name, test := range testTable {
		t.Run(name, func(t *testing.T) {
			article := Article{PublishMonth: test.PublishMonth}
			if article.Quarter() != test.Expected {
				t.Errorf("recieved unexpected result: Got %d, Expected %d", article.Quarter(), test.Expected)
			}
		})
	}
}

// Problem 2
func TestListByTitleAndQuarter(t *testing.T) {
	testTable := map[string]struct {
		Input    []Article
		Expected []string
	}{
		"Empty article list": {
			Input: []Article{},
			Expected: []string{
				"No articles found for Q1",
				"No articles found for Q2",
				"No articles found for Q3",
				"No articles found for Q4",
			},
		},
		"Articles for every quarter": {
			Input: []Article{
				{Title: "A", PublishMonth: time.Date(
					2019, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Title: "B", PublishMonth: time.Date(
					2019, 8, 1, 0, 0, 0, 0, time.UTC)},
				{Title: "C", PublishMonth: time.Date(
					2019, 2, 25, 0, 0, 0, 0, time.UTC)},
				{Title: "D", PublishMonth: time.Date(
					2019, 11, 30, 0, 0, 0, 0, time.UTC)},
				{Title: "E", PublishMonth: time.Date(
					2019, 5, 30, 0, 0, 0, 0, time.UTC)},
			},
			Expected: []string{
				"A - Q1",
				"B - Q3",
				"C - Q1",
				"D - Q4",
				"E - Q2",
			}},
		"Articles for Q1 and Q4 only": {
			Input: []Article{
				{Title: "A", PublishMonth: time.Date(
					2019, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Title: "B", PublishMonth: time.Date(
					2019, 11, 1, 0, 0, 0, 0, time.UTC)},
				{Title: "C", PublishMonth: time.Date(
					2019, 2, 25, 0, 0, 0, 0, time.UTC)},
			},
			Expected: []string{
				"A - Q1",
				"B - Q4",
				"C - Q1",
				"No articles found for Q2",
				"No articles found for Q3",
			},
		},
	}

	for name, test := range testTable {
		t.Run(name, func(t *testing.T) {
			results := ListByTitleAndQuarter(test.Input)
			sort.Strings(results)
			if !cmp.Equal(test.Expected, results) {
				t.Errorf(
					"unexpected response from function. Got [%s], Expected [%s]",
					strings.Join(results, ", "), strings.Join(test.Expected, ", "))
			}
		})
	}
}

// Problem 3
func TestGroupByTopic(t *testing.T) {
	testTable := map[string]struct {
		Articles []Article
		Expected map[string][]Article
	}{
		"No Articles": {Articles: []Article{}, Expected: map[string][]Article{}},
		"Many Articles": {
			Articles: []Article{
				{Id: 1, Title: "A", Topics: []string{"Topic A", "Topic B"}},
				{Id: 2, Title: "B", Topics: []string{"Topic A", "Topic C"}},
				{Id: 3, Title: "C", Topics: []string{"Topic B", "Topic C", "Topic D"}},
				{Id: 4, Title: "D", Topics: []string{"Topic D"}},
				{Id: 5, Title: "E", Topics: []string{"Topic E"}},
			},
			Expected: map[string][]Article{
				"Topic A": {
					{Id: 1, Title: "A", Topics: []string{"Topic A", "Topic B"}},
					{Id: 2, Title: "B", Topics: []string{"Topic A", "Topic C"}},
				},
				"Topic B": {
					{Id: 1, Title: "A", Topics: []string{"Topic A", "Topic B"}},
					{Id: 3, Title: "C", Topics: []string{"Topic B", "Topic C", "Topic D"}},
				},
				"Topic C": {
					{Id: 2, Title: "B", Topics: []string{"Topic A", "Topic C"}},
					{Id: 3, Title: "C", Topics: []string{"Topic B", "Topic C", "Topic D"}},
				},
				"Topic D": {
					{Id: 3, Title: "C", Topics: []string{"Topic B", "Topic C", "Topic D"}},
					{Id: 4, Title: "D", Topics: []string{"Topic D"}},
				},
				"Topic E": {
					{Id: 5, Title: "E", Topics: []string{"Topic E"}},
				},
			},
		},
		"Missing Topics": {
			Articles: []Article{
				{Id: 1, Title: "A", Topics: []string{"Topic A", "Topic B"}},
				{Id: 2, Title: "B", Topics: []string{}},
			},
			Expected: map[string][]Article{
				"Topic A": {
					{Id: 1, Title: "A", Topics: []string{"Topic A", "Topic B"}},
				},
				"Topic B": {
					{Id: 1, Title: "A", Topics: []string{"Topic A", "Topic B"}},
				},
			},
		},
	}

	for name, test := range testTable {
		t.Run(name, func(t *testing.T) {
			results := GroupByTopic(test.Articles)
			for topic, articles := range test.Expected {
				if !cmp.Equal(articles, results[topic]) {
					t.Fail()
				}
			}
		})
	}
}

// Problem 4
func TestGetRelatedTopics(t *testing.T) {
	testTable := map[string]struct {
		Article       Article
		Score         int
		RelatedTopics []string
	}{
		"No relation between articles": {
			Article:       Article{Title: "A", Topics: []string{"Topic D", "Topic E", "Topic F"}},
			Score:         0,
			RelatedTopics: []string{},
		},
		"Several related Topics": {
			Article:       Article{Title: "A", Topics: []string{"Topic A", "Topic B", "Topic C"}},
			Score:         3,
			RelatedTopics: []string{"Topic A", "Topic B", "Topic C"},
		},
	}

	article := &Article{Title: "Control", Topics: []string{"Topic A", "Topic B", "Topic C"}}

	for name, test := range testTable {
		t.Run(name, func(t *testing.T) {
			score, related := article.GetRelatedTopics(&test.Article)
			if score != test.Score {
				t.Errorf("incorrect score given. Got %d, Expected %d", score, test.Score)
			}
			if !cmp.Equal(related, test.RelatedTopics) {
				t.Errorf("incorrect topics returned. Got [%s], Expected [%s]", strings.Join(related, ", "), strings.Join(test.RelatedTopics, ", "))
			}
		})
	}
}

// Problem 5
func TestFindRelatedArticles(t *testing.T) {
	needle := Article{Topics: []string{"Topic A", "Topic C", "Topic E"}}
	testTable := map[string]struct {
		Haystack []Article
		Expected []Article
	}{
		"No articles related": {
			Haystack: []Article{
				{Topics: []string{"Topic B", "Topic F"}},
				{Topics: []string{"Topic D", "Topic F"}},
			},
			Expected: []Article{}},
		"Articles are related": {
			Haystack: []Article{
				{Id: 1, Topics: []string{"Topic A", "Topic C"}},
				{Id: 2, Topics: []string{"Topic A", "Topic E"}},
				{Id: 3, Topics: []string{"Topic E", "Topic B"}},
				{Id: 4, Topics: []string{"Topic D"}},
				{Id: 5, Topics: []string{"Topic B", "Topic C"}},
			},
			Expected: []Article{
				{Id: 1, Topics: []string{"Topic A", "Topic C"}},
				{Id: 2, Topics: []string{"Topic A", "Topic E"}},
				{Id: 3, Topics: []string{"Topic E", "Topic B"}},
				{Id: 5, Topics: []string{"Topic B", "Topic C"}},
			},
		},
		"Related articles are sorted by weight": {
			Haystack: []Article{
				{Id: 1, Topics: []string{"Topic A", "Topic C", "Topic E"}},
				{Id: 2, Topics: []string{"Topic A", "Topic B", "Topic D", "Topic F"}},
				{Id: 3, Topics: []string{"Topic A", "Topic C"}},
			},
			Expected: []Article{
				{Id: 1, Topics: []string{"Topic A", "Topic C", "Topic E"}},
				{Id: 3, Topics: []string{"Topic A", "Topic C"}},
				{Id: 2, Topics: []string{"Topic A", "Topic B", "Topic D", "Topic F"}},
			},
		},
	}

	for name, test := range testTable {
		t.Run(name, func(t *testing.T) {
			related := FindRelatedArticles(needle, test.Haystack)
			if !cmp.Equal(related, test.Expected) {
				t.Fail()
			}
		})
	}
}