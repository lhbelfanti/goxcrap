package search

type (
	// CriteriaDTO represents a slice of CriterionDTO
	CriteriaDTO []CriterionDTO

	// CriterionDTO is a Data Transfer Object used to represent a Criterion type
	CriterionDTO struct {
		ID               string   `json:"id,omitempty"`
		AllOfTheseWords  []string `json:"all_of_these_words,omitempty"`
		ThisExactPhrase  string   `json:"this_exact_phrase,omitempty"`
		AnyOfTheseWords  []string `json:"any_of_these_words,omitempty"`
		NoneOfTheseWords []string `json:"none_of_these_words,omitempty"`
		TheseHashtags    []string `json:"these_hashtags,omitempty"`
		Language         string   `json:"language,omitempty"`
		Since            string   `json:"since,omitempty"`
		Until            string   `json:"until,omitempty"`
	}
)

// ToType converts a CriteriaDTO into a Criteria
func (c CriteriaDTO) ToType() Criteria {
	criteria := make(Criteria, 0, len(c))
	for _, criterion := range c {
		criteria = append(criteria, criterion.ToType())
	}

	return criteria
}

// ToType converts a CriterionDTO into a Criterion
func (c CriterionDTO) ToType() Criterion {
	return Criterion{
		ID:               c.ID,
		AllOfTheseWords:  c.AllOfTheseWords,
		ThisExactPhrase:  c.ThisExactPhrase,
		AnyOfTheseWords:  c.AnyOfTheseWords,
		NoneOfTheseWords: c.NoneOfTheseWords,
		TheseHashtags:    c.TheseHashtags,
		Language:         c.Language,
		Since:            c.Since,
		Until:            c.Until,
	}
}
