package tzkit

import "testing"

func TestGetTimezone(t *testing.T) {
	tests := []struct {
		name       string
		zoneID     string
		wantFound  bool
		wantHasDST bool
	}{
		{
			name:       "Valid zone with DST",
			zoneID:     "America/New_York",
			wantFound:  true,
			wantHasDST: true,
		},
		{
			name:       "Valid zone without DST",
			zoneID:     "Asia/Tokyo",
			wantFound:  true,
			wantHasDST: false,
		},
		{
			name:      "Invalid zone",
			zoneID:    "Fake/Zone",
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz, ok := GetTimezone(tt.zoneID)

			if ok != tt.wantFound {
				t.Fatalf("GetTimezone() got ok = %v, want %v", ok, tt.wantFound)
			}

			if ok {
				if tz.ID != tt.zoneID {
					t.Errorf("Expected ID %s, got %s", tt.zoneID, tz.ID)
				}
				if tz.HasDST() != tt.wantHasDST {
					t.Errorf("HasDST() got %v, want %v", tz.HasDST(), tt.wantHasDST)
				}
			}
		})
	}
}

func TestTimezoneIDs(t *testing.T) {
	ids := TimezoneIDs()
	if len(ids) == 0 {
		t.Fatal("expected at least one timezone ID")
	}
	for i := 1; i < len(ids); i++ {
		if ids[i] < ids[i-1] {
			t.Fatalf("expected sorted IDs, but %s came after %s", ids[i], ids[i-1])
		}
	}
}

func TestTimezonesByCountry(t *testing.T) {
	tests := []struct {
		name        string
		countryCode string
		wantEmpty   bool
	}{
		{
			name:        "Valid country with timezones",
			countryCode: "US",
			wantEmpty:   false,
		},
		{
			name:        "Invalid country code",
			countryCode: "ZZ",
			wantEmpty:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tzs := TimezonesByCountry(tt.countryCode)

			if tt.wantEmpty && len(tzs) != 0 {
				t.Fatalf("expected no timezones, got %d", len(tzs))
			}
			if !tt.wantEmpty && len(tzs) == 0 {
				t.Fatal("expected at least one timezone")
			}
		})
	}
}
