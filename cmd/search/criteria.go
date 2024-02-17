package search

type GetAdvanceSearchCriteria func() []Criteria

// MakeGetAdvanceSearchCriteria creates a new GetAdvanceSearchCriteria
func MakeGetAdvanceSearchCriteria() GetAdvanceSearchCriteria {
	return func() []Criteria {
		return []Criteria{
			{
				ID:               "Example",
				AllOfTheseWords:  []string{"casa", "perro"},
				ThisExactPhrase:  "mi mascota",
				AnyOfTheseWords:  []string{"cachorro"},
				NoneOfTheseWords: []string{"gato"},
				TheseHashtags:    []string{"#mascotas", "#canino"},
				Language:         "es",
				Since:            "2006-01-01",
				Until:            "2024-01-01",
			},
		}
	}
}
