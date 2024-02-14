package search_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/search"
)

func TestParseDate(t *testing.T) {
	for _, test := range []struct {
		date string
		want search.Date // if empty, expect an error
	}{
		{"2023-01-02", search.Date{Year: 2023, Month: 1, Day: 2}},
		{"2023-12-31", search.Date{Year: 2023, Month: 12, Day: 31}},
		{"0003-02-04", search.Date{Year: 3, Month: 2, Day: 4}},
		{"999-01-26", search.Date{}},
		{"", search.Date{}},
		{"2023-01-02x", search.Date{}},
	} {
		got, err := search.ParseDate(test.date)
		if err != nil {
			assert.Equal(t, got, search.Date{})
		} else {
			assert.Equal(t, got, test.want)
		}
	}
}

func TestCriteria_ParseDates(t *testing.T) {
	for _, test := range []struct {
		criteria search.Criteria
		want1    search.Date // if empty, expect an error
		want2    search.Date // if empty, expect an error
	}{
		{
			search.Criteria{Since: "2023-01-02", Until: "2023-01-03"},
			search.Date{Year: 2023, Month: 1, Day: 2},
			search.Date{Year: 2023, Month: 1, Day: 3},
		},
		{
			search.Criteria{Since: "2023-12-31", Until: "2024-01-31"},
			search.Date{Year: 2023, Month: 12, Day: 31},
			search.Date{Year: 2024, Month: 1, Day: 31},
		},
		{
			search.Criteria{Since: "0003-12-31", Until: "0004-01-31"},
			search.Date{Year: 3, Month: 12, Day: 31},
			search.Date{Year: 4, Month: 1, Day: 31},
		},
		{search.Criteria{Since: "999-01-26", Until: "999-01-28"}, search.Date{}, search.Date{}},
		{search.Criteria{Since: "2024-01-26", Until: "999-01-28"}, search.Date{}, search.Date{}},
		{search.Criteria{Since: "", Until: "2024-01-28"}, search.Date{}, search.Date{}},
		{search.Criteria{Since: "2024-01-26", Until: ""}, search.Date{}, search.Date{}},
		{search.Criteria{Since: "2024-01-26e", Until: "2024-01-28"}, search.Date{}, search.Date{}},
		{search.Criteria{Since: "2024-01-26", Until: "2024-01-28e"}, search.Date{}, search.Date{}},
	} {
		got1, got2, err := test.criteria.ParseDates()
		if err != nil {
			assert.Equal(t, got1, search.Date{})
			assert.Equal(t, got2, search.Date{})
		} else {
			assert.Equal(t, got1, test.want1)
			assert.Equal(t, got2, test.want2)
		}
	}
}

func TestDate_AddDays(t *testing.T) {
	for _, test := range []struct {
		desc  string
		start search.Date
		end   search.Date
		days  int
	}{
		{
			desc:  "zero days noop",
			start: search.Date{Year: 2014, Month: 5, Day: 9},
			end:   search.Date{Year: 2014, Month: 5, Day: 9},
			days:  0,
		},
		{
			desc:  "crossing a year boundary",
			start: search.Date{Year: 2014, Month: 12, Day: 31},
			end:   search.Date{Year: 2015, Month: 1, Day: 1},
			days:  1,
		},
		{
			desc:  "negative number of days",
			start: search.Date{Year: 2015, Month: 1, Day: 1},
			end:   search.Date{Year: 2014, Month: 12, Day: 31},
			days:  -1,
		},
		{
			desc:  "full leap year",
			start: search.Date{Year: 2004, Month: 1, Day: 1},
			end:   search.Date{Year: 2005, Month: 1, Day: 1},
			days:  366,
		},
		{
			desc:  "full non-leap year",
			start: search.Date{Year: 2001, Month: 1, Day: 1},
			end:   search.Date{Year: 2002, Month: 1, Day: 1},
			days:  365,
		},
		{
			desc:  "crossing a leap second",
			start: search.Date{Year: 1972, Month: 6, Day: 30},
			end:   search.Date{Year: 1972, Month: 7, Day: 1},
			days:  1,
		},
		{
			desc:  "dates before the unix epoch",
			start: search.Date{Year: 101, Month: 1, Day: 1},
			end:   search.Date{Year: 102, Month: 1, Day: 1},
			days:  365,
		},
	} {
		got := test.start.AddDays(test.days)
		assert.Equal(t, got, test.end)
	}
}

func TestDate_String(t *testing.T) {
	for _, test := range []struct {
		date search.Date
		want string
	}{
		{search.Date{Year: 2023, Month: 1, Day: 2}, "2023-01-02"},
		{search.Date{Year: 2023, Month: 12, Day: 31}, "2023-12-31"},
		{search.Date{Year: 3, Month: 2, Day: 4}, "0003-02-04"},
		{search.Date{Year: 999, Month: 1, Day: 26}, "0999-01-26"},
		{search.Date{}, "0000-00-00"},
	} {
		got := test.date.String()
		assert.Equal(t, got, test.want)
	}
}
