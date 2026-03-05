package tzkit

import "slices"

// Timezone represents an IANA timezone with its offsets and abbreviations.
type Timezone struct {
	ID              string   `json:"id"`
	CountryCodes    []string `json:"country_codes"`
	Type            string   `json:"type"`
	SDTOffset       int      `json:"sdt_offset"`
	DSTOffset       int      `json:"dst_offset"`
	SDTAbbreviation string   `json:"sdt_abbreviation"`
	DSTAbbreviation string   `json:"dst_abbreviation"`
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
