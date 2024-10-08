package criteria

import "encoding/json"

type (
	// DTO is a Data Transfer Object used to represent a Type
	DTO struct {
		ID               int      `json:"id"`
		Name             string   `json:"name,omitempty"`
		AllOfTheseWords  []string `json:"all_of_these_words,omitempty"`
		ThisExactPhrase  string   `json:"this_exact_phrase,omitempty"`
		AnyOfTheseWords  []string `json:"any_of_these_words,omitempty"`
		NoneOfTheseWords []string `json:"none_of_these_words,omitempty"`
		TheseHashtags    []string `json:"these_hashtags,omitempty"`
		Language         string   `json:"language,omitempty"`
		Since            string   `json:"since,omitempty"`
		Until            string   `json:"until,omitempty"`
	}

	// IncomingBrokerMessageDTO is the message to enqueue in the message broker
	IncomingBrokerMessageDTO struct {
		Message json.RawMessage `json:"message"`
	}

	// MessageDTO is the necessary struct to decode the message consumed from the broker
	MessageDTO struct {
		Criteria    DTO `json:"criteria"`
		ExecutionID int `json:"execution_id"`
	}
)

// ToType converts a DTO into a Type
func (c DTO) ToType() Type {
	return Type{
		ID:               c.ID,
		Name:             c.Name,
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
