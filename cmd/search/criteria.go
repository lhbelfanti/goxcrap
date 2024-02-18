package search

type GetAdvanceSearchCriteria func() []Criteria

// MakeGetAdvanceSearchCriteria creates a new GetAdvanceSearchCriteria
func MakeGetAdvanceSearchCriteria() GetAdvanceSearchCriteria {
	return func() []Criteria {
		return []Criteria{
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
}
