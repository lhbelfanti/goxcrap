package search

type GetSearchCriteria func() []Criteria

// MakeGetSearchCriteria creates a new GetSearchCriteria
func MakeGetSearchCriteria() GetSearchCriteria {
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
