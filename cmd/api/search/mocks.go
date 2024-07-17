package search

// MockCriteria mocks a slice of Criterion
func MockCriteria() Criteria {
	return Criteria{
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

// MockCriteriaDTO mocks a slice of CriterionDTO
func MockCriteriaDTO() CriteriaDTO {
	return CriteriaDTO{
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

// MockExampleCriteria mocks a real example of a slice of Criterion. Only for local and test use
func MockExampleCriteria() Criteria {
	return Criteria{
		{
			ID:               "Example",
			AllOfTheseWords:  []string{"casa"},
			ThisExactPhrase:  "",
			AnyOfTheseWords:  []string{"perro"},
			NoneOfTheseWords: []string{"gato"},
			TheseHashtags:    nil,
			Language:         "es",
			Since:            "2024-01-01",
			Until:            "2024-01-02",
		},
	}
}

// MockExecuteAdvanceSearch mocks ExecuteAdvanceSearch function
func MockExecuteAdvanceSearch(err error) ExecuteAdvanceSearch {
	return func(criterion Criterion) error {
		return err
	}
}
