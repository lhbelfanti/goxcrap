package tweets

type RetrieveAll func() []Tweet

func MakeRetrieveAll() RetrieveAll {
	return func() []Tweet {
		return []Tweet{
			{
				Text:   "Testing test",
				Links:  nil,
				Images: nil,
			},
		}
	}
}
