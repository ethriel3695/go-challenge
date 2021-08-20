SHELL := /bin/bash

#
# Local Specific Commands
#
test-all: TestQuarterGetter TestListByTitleAndQuarter TestGroupByTopic TestGetRelatedTopics TestFindRelatedArticles

# Problem 1
TestQuarterGetter:
	go test -v -run TestQuarterGetter

# Problem 2
TestListByTitleAndQuarter:
	go test -v -run TestListByTitleAndQuarter

# Problem 3
TestGroupByTopic:
	go test -v -run TestGroupByTopic

# Problem 4
TestGetRelatedTopics:
	go test -v -run TestGroupByTopic

# Problem 5
TestFindRelatedArticles:
	go test -v -run TestFindRelatedArticles