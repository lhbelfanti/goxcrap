package search

// MockMakeGetSearchCriteria mocks the function MakeGetSearchCriteria and the values returned by GetSearchCriteria
func MockMakeGetSearchCriteria(criteria []Criteria) GetSearchCriteria {
	return func() []Criteria {
		return criteria
	}
}

// MockCriteria mocks a slice of Criteria
func MockCriteria() []Criteria {
	return []Criteria{
		{
			ID:               "Example",
			AllOfTheseWords:  []string{"word1", "word2"},
			ThisExactPhrase:  "exact phrase",
			AnyOfTheseWords:  []string{"any1", "any2"},
			NoneOfTheseWords: []string{"none1", "none2"},
			TheseHashtags:    []string{"#hashtag1", "#hashtag2"},
			Language:         "es",
			Since:            "2006-01-01",
			Until:            "2024-01-01",
		},
	}
}
