package tzkit

import "testing"

func TestGetCountry(t *testing.T) {
	tests := []struct {
		name      string
		code      string
		wantFound bool
		wantName  string
	}{
		{
			name:      "Valid country",
			code:      "US",
			wantFound: true,
			wantName:  "United States",
		},
		{
			name:      "Another valid country",
			code:      "JP",
			wantFound: true,
			wantName:  "Japan",
		},
		{
			name:      "Invalid country code",
			code:      "ZZ",
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			country, ok := GetCountry(tt.code)

			if ok != tt.wantFound {
				t.Fatalf("GetCountry() got ok = %v, want %v", ok, tt.wantFound)
			}

			if ok && country.Name != tt.wantName {
				t.Errorf("Expected name %s, got %s", tt.wantName, country.Name)
			}
		})
	}
}

func TestCountryCodes(t *testing.T) {
	codes := CountryCodes()
	if len(codes) == 0 {
		t.Fatal("expected at least one country code")
	}
	for i := 1; i < len(codes); i++ {
		if codes[i] < codes[i-1] {
			t.Fatalf("expected sorted codes, but %s came after %s", codes[i], codes[i-1])
		}
	}
}

func TestCountriesByTimezone(t *testing.T) {
	tests := []struct {
		name      string
		tzID      string
		wantEmpty bool
	}{
		{
			name:      "Valid timezone",
			tzID:      "America/New_York",
			wantEmpty: false,
		},
		{
			name:      "Invalid timezone",
			tzID:      "Fake/Zone",
			wantEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countries := CountriesByTimezone(tt.tzID)

			if tt.wantEmpty && len(countries) != 0 {
				t.Fatalf("expected no countries, got %d", len(countries))
			}
			if !tt.wantEmpty && len(countries) == 0 {
				t.Fatal("expected at least one country")
			}
		})
	}
}
