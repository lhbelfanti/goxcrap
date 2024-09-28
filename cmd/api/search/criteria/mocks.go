package criteria

// MockCriteria mocks a criteria.Type
func MockCriteria() Type {
	return Type{
		ID:               1,
		Name:             "Example",
		AllOfTheseWords:  []string{"word1", "word2"},
		ThisExactPhrase:  "exact phrase",
		AnyOfTheseWords:  []string{"any1", "any2"},
		NoneOfTheseWords: []string{"none1", "none2"},
		TheseHashtags:    []string{"#hashtag1", "#hashtag2"},
		Language:         "es",
		Since:            "2006-01-01",
		Until:            "2024-01-01",
	}
}

// MockCriteriaDTO mocks a criteria.DTO
func MockCriteriaDTO() DTO {
	return DTO{
		ID:               1,
		Name:             "Example",
		AllOfTheseWords:  []string{"word1", "word2"},
		ThisExactPhrase:  "exact phrase",
		AnyOfTheseWords:  []string{"any1", "any2"},
		NoneOfTheseWords: []string{"none1", "none2"},
		TheseHashtags:    []string{"#hashtag1", "#hashtag2"},
		Language:         "es",
		Since:            "2006-01-01",
		Until:            "2024-01-01",
	}
}

// MockExampleCriteria mocks a real example of a criteria.Type. Only for local and test purposes
func MockExampleCriteria() Type {
	return Type{
		ID:               1,
		Name:             "Example",
		AllOfTheseWords:  []string{"casa"},
		ThisExactPhrase:  "",
		AnyOfTheseWords:  []string{"perro"},
		NoneOfTheseWords: []string{"gato"},
		TheseHashtags:    nil,
		Language:         "es",
		Since:            "2024-01-01",
		Until:            "2024-01-02",
	}
}

// MockMessageDTO mocks a criteria.MessageDTO
func MockMessageDTO() MessageDTO {
	return MessageDTO{
		Criteria:    MockCriteriaDTO(),
		ExecutionID: 1,
	}
}
