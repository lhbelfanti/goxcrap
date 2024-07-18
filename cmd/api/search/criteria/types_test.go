package criteria_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/search/criteria"
)

func TestParseDate(t *testing.T) {
	for _, test := range []struct {
		date string
		want criteria.Date // if empty, expect an error
	}{
		{"2023-01-02", criteria.Date{Year: 2023, Month: 1, Day: 2}},
		{"2023-12-31", criteria.Date{Year: 2023, Month: 12, Day: 31}},
		{"0003-02-04", criteria.Date{Year: 3, Month: 2, Day: 4}},
		{"999-01-26", criteria.Date{}},
		{"", criteria.Date{}},
		{"2023-01-02x", criteria.Date{}},
	} {
		got, err := criteria.ParseDate(test.date)
		if err != nil {
			assert.Equal(t, criteria.Date{}, got)
		} else {
			assert.Equal(t, test.want, got)
		}
	}
}

func TestDate_String(t *testing.T) {
	for _, test := range []struct {
		date criteria.Date
		want string
	}{
		{criteria.Date{Year: 2023, Month: 1, Day: 2}, "2023-01-02"},
		{criteria.Date{Year: 2023, Month: 12, Day: 31}, "2023-12-31"},
		{criteria.Date{Year: 3, Month: 2, Day: 4}, "0003-02-04"},
		{criteria.Date{Year: 999, Month: 1, Day: 26}, "0999-01-26"},
		{criteria.Date{}, "0000-00-00"},
	} {
		got := test.date.String()
		assert.Equal(t, test.want, got)
	}
}

func TestDate_AddDays(t *testing.T) {
	for _, test := range []struct {
		desc  string
		start criteria.Date
		end   criteria.Date
		days  int
	}{
		{
			desc:  "zero days noop",
			start: criteria.Date{Year: 2014, Month: 5, Day: 9},
			end:   criteria.Date{Year: 2014, Month: 5, Day: 9},
			days:  0,
		},
		{
			desc:  "crossing a year boundary",
			start: criteria.Date{Year: 2014, Month: 12, Day: 31},
			end:   criteria.Date{Year: 2015, Month: 1, Day: 1},
			days:  1,
		},
		{
			desc:  "negative number of days",
			start: criteria.Date{Year: 2015, Month: 1, Day: 1},
			end:   criteria.Date{Year: 2014, Month: 12, Day: 31},
			days:  -1,
		},
		{
			desc:  "full leap year",
			start: criteria.Date{Year: 2004, Month: 1, Day: 1},
			end:   criteria.Date{Year: 2005, Month: 1, Day: 1},
			days:  366,
		},
		{
			desc:  "full non-leap year",
			start: criteria.Date{Year: 2001, Month: 1, Day: 1},
			end:   criteria.Date{Year: 2002, Month: 1, Day: 1},
			days:  365,
		},
		{
			desc:  "crossing a leap second",
			start: criteria.Date{Year: 1972, Month: 6, Day: 30},
			end:   criteria.Date{Year: 1972, Month: 7, Day: 1},
			days:  1,
		},
		{
			desc:  "dates before the unix epoch",
			start: criteria.Date{Year: 101, Month: 1, Day: 1},
			end:   criteria.Date{Year: 102, Month: 1, Day: 1},
			days:  365,
		},
	} {
		got := test.start.AddDays(test.days)
		assert.Equal(t, test.end, got)
	}
}

func TestDate_After(t *testing.T) {
	for _, test := range []struct {
		date1 criteria.Date
		date2 criteria.Date
		want  bool
	}{
		{criteria.Date{Year: 2024, Month: 2, Day: 15}, criteria.Date{Year: 2024, Month: 2, Day: 14}, true},
		{criteria.Date{Year: 2024, Month: 2, Day: 14}, criteria.Date{Year: 2024, Month: 2, Day: 15}, false},
		{criteria.Date{Year: 2024, Month: 2, Day: 15}, criteria.Date{Year: 2024, Month: 2, Day: 15}, false},
		{criteria.Date{Year: 2024, Month: 3, Day: 14}, criteria.Date{Year: 2024, Month: 2, Day: 14}, true},
		{criteria.Date{Year: 2024, Month: 2, Day: 14}, criteria.Date{Year: 2024, Month: 3, Day: 15}, false},
		{criteria.Date{Year: 2025, Month: 2, Day: 15}, criteria.Date{Year: 2024, Month: 2, Day: 14}, true},
		{criteria.Date{Year: 2024, Month: 2, Day: 14}, criteria.Date{Year: 2025, Month: 2, Day: 15}, false},
	} {
		got := test.date1.After(test.date2)
		assert.Equal(t, test.want, got)
	}
}

func TestCriteria_ParseDates(t *testing.T) {
	for _, test := range []struct {
		searchCriteria criteria.Type
		want1          criteria.Date // if empty, expect an error
		want2          criteria.Date // if empty, expect an error
	}{
		{
			criteria.Type{Since: "2023-01-02", Until: "2023-01-03"},
			criteria.Date{Year: 2023, Month: 1, Day: 2},
			criteria.Date{Year: 2023, Month: 1, Day: 3},
		},
		{
			criteria.Type{Since: "2023-12-31", Until: "2024-01-31"},
			criteria.Date{Year: 2023, Month: 12, Day: 31},
			criteria.Date{Year: 2024, Month: 1, Day: 31},
		},
		{
			criteria.Type{Since: "0003-12-31", Until: "0004-01-31"},
			criteria.Date{Year: 3, Month: 12, Day: 31},
			criteria.Date{Year: 4, Month: 1, Day: 31},
		},
		{criteria.Type{Since: "999-01-26", Until: "999-01-28"}, criteria.Date{}, criteria.Date{}},
		{criteria.Type{Since: "2024-01-26", Until: "999-01-28"}, criteria.Date{}, criteria.Date{}},
		{criteria.Type{Since: "", Until: "2024-01-28"}, criteria.Date{}, criteria.Date{}},
		{criteria.Type{Since: "2024-01-26", Until: ""}, criteria.Date{}, criteria.Date{}},
		{criteria.Type{Since: "2024-01-26e", Until: "2024-01-28"}, criteria.Date{}, criteria.Date{}},
		{criteria.Type{Since: "2024-01-26", Until: "2024-01-28e"}, criteria.Date{}, criteria.Date{}},
	} {
		got1, got2, err := test.searchCriteria.ParseDates()
		if err != nil {
			assert.Equal(t, criteria.Date{}, got1)
			assert.Equal(t, criteria.Date{}, got2)
		} else {
			assert.Equal(t, test.want1, got1)
			assert.Equal(t, test.want2, got2)
		}
	}
}

func TestCriteria_ConvertIntoQueryString(t *testing.T) {
	for _, test := range []struct {
		searchCriteria criteria.Type
		want           string
	}{
		{
			searchCriteria: criteria.Type{
				AllOfTheseWords:  []string{"all1", "all2"},
				ThisExactPhrase:  "exact phrase",
				AnyOfTheseWords:  []string{"any1", "any2"},
				NoneOfTheseWords: []string{"none1", "none2"},
				TheseHashtags:    []string{"#hashtag1", "#hashtag2"},
				Language:         "es",
				Since:            "2006-01-01",
				Until:            "2024-01-01",
			},
			want: "q=all1%20all2%20%22exact%20phrase%22%20(any1%20OR%20any2)%20-none1%20-none2%20(%23hashtag1%20OR%20%23hashtag2)%20lang:es%20until:2024-01-01%20since:2006-01-01&src=typed_query",
		},
		{
			searchCriteria: criteria.Type{
				ThisExactPhrase:  "exact phrase",
				AnyOfTheseWords:  []string{"any1", "any2"},
				NoneOfTheseWords: []string{"none1", "none2"},
				TheseHashtags:    []string{"#hashtag1", "#hashtag2"},
				Language:         "es",
				Since:            "2006-01-01",
				Until:            "2024-01-01",
			},
			want: "q=%22exact%20phrase%22%20(any1%20OR%20any2)%20-none1%20-none2%20(%23hashtag1%20OR%20%23hashtag2)%20lang:es%20until:2024-01-01%20since:2006-01-01&src=typed_query",
		},
		{
			searchCriteria: criteria.Type{
				AnyOfTheseWords:  []string{"any1", "any2"},
				NoneOfTheseWords: []string{"none1", "none2"},
				TheseHashtags:    []string{"#hashtag1", "#hashtag2"},
				Language:         "es",
				Since:            "2006-01-01",
				Until:            "2024-01-01",
			},
			want: "q=(any1%20OR%20any2)%20-none1%20-none2%20(%23hashtag1%20OR%20%23hashtag2)%20lang:es%20until:2024-01-01%20since:2006-01-01&src=typed_query",
		},
		{
			searchCriteria: criteria.Type{
				NoneOfTheseWords: []string{"none1", "none2"},
				TheseHashtags:    []string{"#hashtag1", "#hashtag2"},
				Language:         "es",
				Since:            "2006-01-01",
				Until:            "2024-01-01",
			},
			want: "q=-none1%20-none2%20(%23hashtag1%20OR%20%23hashtag2)%20lang:es%20until:2024-01-01%20since:2006-01-01&src=typed_query",
		},
		{
			searchCriteria: criteria.Type{
				TheseHashtags: []string{"#hashtag1", "#hashtag2"},
				Language:      "es",
				Since:         "2006-01-01",
				Until:         "2024-01-01",
			},
			want: "q=(%23hashtag1%20OR%20%23hashtag2)%20lang:es%20until:2024-01-01%20since:2006-01-01&src=typed_query",
		},
		{
			searchCriteria: criteria.Type{
				Language: "es",
				Since:    "2006-01-01",
				Until:    "2024-01-01",
			},
			want: "q=lang:es%20until:2024-01-01%20since:2006-01-01&src=typed_query",
		},
		{
			searchCriteria: criteria.Type{
				Since: "2006-01-01",
				Until: "2024-01-01",
			},
			want: "q=until:2024-01-01%20since:2006-01-01&src=typed_query",
		},
		{
			searchCriteria: criteria.Type{
				Until: "2024-01-01",
			},
			want: "q=until:2024-01-01&src=typed_query",
		},
		{
			searchCriteria: criteria.Type{},
			want:           "q=&src=typed_query",
		},
	} {
		got := test.searchCriteria.ConvertIntoQueryString()
		assert.Equal(t, test.want, got)
	}
}
