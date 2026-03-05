package tzkit

import (
	"fmt"
	"slices"
	"time"
)

// Timezone represents an IANA timezone with its offsets and abbreviations.
type Timezone struct {
	ID              string   `json:"id"`
	CountryCodes    []string `json:"country_codes"`
	Type            string   `json:"type"`
	SDTOffset       int      `json:"sdt_offset"`
	DSTOffset       int      `json:"dst_offset"`
	SDTAbbreviation string   `json:"sdt_abbreviation"`
	DSTAbbreviation string   `json:"dst_abbreviation"`
	SDTName         string   `json:"sdt_name"`
	DSTName         string   `json:"dst_name"`
	Link            string   `json:"link"`
}

// GetTimezone returns the Timezone for the given IANA timezone ID (e.g. "America/New_York").
// The second return value reports whether the ID was found.
func GetTimezone(id string) (Timezone, bool) {
	tz, ok := Timezones[id]
	return tz, ok
}

// TimezoneIDs returns a sorted slice of all IANA timezone IDs.
func TimezoneIDs() []string {
	ids := make([]string, 0, len(Timezones))
	for id := range Timezones {
		ids = append(ids, id)
	}
	slices.Sort(ids)
	return ids
}

// TimezonesByCountry returns all timezones associated with the given ISO 3166-1 alpha-2 country code (e.g. "US").
func TimezonesByCountry(countryCode string) []Timezone {
	var result []Timezone
	for _, tz := range Timezones {
		if slices.Contains(tz.CountryCodes, countryCode) {
			result = append(result, tz)
		}
	}
	return result
}

// HasDST reports whether the timezone observes daylight saving time.
func (tz Timezone) HasDST() bool {
	return tz.SDTOffset != tz.DSTOffset
}

// OffsetMinutes returns the UTC offset in minutes for the given time,
// accounting for whether DST is active.
func (tz Timezone) OffsetMinutes(t time.Time) int {
	if tz.HasDST() && tz.In(t).IsDST() {
		return tz.DSTOffset
	}
	return tz.SDTOffset
}

// In converts a time.Time to this timezone.
func (tz Timezone) In(t time.Time) time.Time {
	loc, err := time.LoadLocation(tz.ID)
	if err != nil {
		return t
	}
	return t.In(loc)
}

// FormatTime formats a time.Time in this timezone like:
// "Thu Mar 05 2026 11:50:27 GMT-0700 (Mountain Standard Time)"
func (tz Timezone) FormatTime(t time.Time) string {
	loc, err := time.LoadLocation(tz.ID)
	if err != nil {
		loc = time.UTC
	}
	t = t.In(loc)

	_, offset := t.Zone()
	hours := offset / 3600
	mins := (offset % 3600) / 60
	if mins < 0 {
		mins = -mins
	}

	name := tz.SDTName
	if tz.HasDST() && t.IsDST() {
		name = tz.DSTName
	}

	return fmt.Sprintf(
		"%s GMT%+03d%02d (%s)",
		t.Format("Mon Jan 02 2006 15:04:05"),
		hours, mins, name,
	)
}
