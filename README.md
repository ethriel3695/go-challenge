# Articles Package

## Purpose for this package

This is a small package created to provide some commonly used methods
useful when dealing with articles from a CMS (content management system).

## Overview

Below are a series of exercises where you must implement a method behavior
in order to pass a unit test.

Please do not change any file names, class names, or method names. The tests and scripts
included are looking for those specific names.. Thanks!

## Problem 1

Implement the `Quarter` method on the `Article` struct.

Run the script `make TestQuarterGetter` to see the test results. Modify `article.go`
in order to make the test pass.You can view the tests in `articles_test.go`.

The test assertions are in the function `TestQuarterGetter`.

## Problem 2

Implement the `ListByTitleAndQuarter` function.

Run the script `make TestListByTitleAndQuarter` to see the test results. Modify `article.go`
in order to make the test pass.You can view the tests in `articles_test.go`.

The test assertions are in the function `TestListByTitleAndQuarter`.

## Problem 3

Implement the `GroupByTopic` function.

Run the script `make TestGroupByTopic` to see the failure. Modify `articles.go`
in order to make the test pass. You can view the tests in `articles_test.go`.

The test assertions are in the function `TestGroupByTopic`.

## Problem 4

Implement the `GetRelatedTopics` method.

Run the script `make TestGetRelatedTopics` to see the failure. Modify `articles.go`
in order to make the test pass. You can view the tests in `articles_test.go`.

The test assertions are in the function `TestGetRelatedTopics`.

## Problem 5

Implement the `FindRelatedArticles` function.

Run the script `make TestFindRelatedArticles` to see the failure. Modify `articles.go`
in order to make the test pass. You can view the tests in `articles_test.go`.

The test assertions are in the method `TestFindRelatedArticles`.