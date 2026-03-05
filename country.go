package tzkit

import "slices"

// Country represents a country with its display name and associated IANA timezone IDs.
type Country struct {
	Name  string   `json:"name"`
	Zones []string `json:"zones"`
}

// GetCountry returns the Country for the given ISO 3166-1 alpha-2 code (e.g. "US").
// The second return value reports whether the code was found.
func GetCountry(code string) (Country, bool) {
	country, ok := Countries[code]
	return country, ok
}

// CountryCodes returns a sorted slice of all ISO 3166-1 alpha-2 country codes.
func CountryCodes() []string {
	codes := make([]string, 0, len(Countries))
	for code := range Countries {
		codes = append(codes, code)
	}
	slices.Sort(codes)
	return codes
}

// CountriesByTimezone returns all countries that use the given IANA timezone ID (e.g. "America/New_York").
func CountriesByTimezone(tzID string) []Country {
	var result []Country
	for _, country := range Countries {
		if slices.Contains(country.Zones, tzID) {
			result = append(result, country)
		}
	}
	return result
}
