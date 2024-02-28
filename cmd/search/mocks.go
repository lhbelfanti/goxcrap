package search

// MockGetAdvanceSearchCriteria mocks GetAdvanceSearchCriteria function
func MockGetAdvanceSearchCriteria(criteria []Criteria) GetAdvanceSearchCriteria {
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

// MockExecuteAdvanceSearch mocks ExecuteAdvanceSearch function
func MockExecuteAdvanceSearch(err error) ExecuteAdvanceSearch {
	return func(searchCriteria Criteria) error {
		return err
	}
}
