package search

import (
	"fmt"
	"time"
)

type (
	// Criteria represents all the criteria that need to be taken in consideration for the tweets search
	Criteria struct {
		ID               string
		AllOfTheseWords  []string
		ThisExactPhrase  []string
		AnyOfTheseWords  []string
		NoneOfTheseWords []string
		TheseHashtags    []string
		Language         string
		Since            string
		Until            string
	}

	// Date represents a Date with a Year, Month and Day
	Date struct {
		Year  int        // Year (e.g., 2014).
		Month time.Month // Month of the year (January = 1, ...).
		Day   int        // Day of the month, starting at 1.
	}
)

// ParseDate parses a string in RFC3339 full-date format and returns the date value it represents
func ParseDate(s string) (Date, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return Date{}, err
	}
	return Of(t), nil
}

// Of returns the Date in which a time occurs in that time's location
func Of(t time.Time) Date {
	var d Date
	d.Year, d.Month, d.Day = t.Date()
	return d
}

// ParseDates parses both Since and Until strings in RFC3339 full-date format and returns the date value it represents
func (c Criteria) ParseDates() (Date, Date, error) {
	since, err := ParseDate(c.Since)
	if err != nil {
		return Date{}, Date{}, err
	}

	until, err := ParseDate(c.Until)
	if err != nil {
		return Date{}, Date{}, err
	}

	return since, until, nil
}

// In returns the time corresponding to time 00:00:00 of the date in the location
func (d Date) In(loc *time.Location) time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, loc)
}

// String returns the date in RFC3339 full-date format.
func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

// AddDays returns the date that is n days in the future. n can also be negative to go into the past
func (d Date) AddDays(n int) Date {
	return Of(d.In(time.UTC).AddDate(0, 0, n))
}
