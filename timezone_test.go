package tzkit

import (
	"testing"
	"time"
)

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

func TestIn(t *testing.T) {
	utc := time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		zone     string
		wantLoc  string
		wantHour int
	}{
		{
			name:     "UTC to Denver",
			zone:     "America/Denver",
			wantLoc:  "America/Denver",
			wantHour: 5,
		},
		{
			name:     "UTC to Tokyo",
			zone:     "Asia/Tokyo",
			wantLoc:  "Asia/Tokyo",
			wantHour: 21,
		},
		{
			name:     "UTC to New York",
			zone:     "America/New_York",
			wantLoc:  "America/New_York",
			wantHour: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz, _ := GetTimezone(tt.zone)
			got := tz.In(utc)

			if got.Location().String() != tt.wantLoc {
				t.Errorf("In() location = %q, want %q", got.Location().String(), tt.wantLoc)
			}
			if got.Hour() != tt.wantHour {
				t.Errorf("In() hour = %d, want %d", got.Hour(), tt.wantHour)
			}
			if !got.Equal(utc) {
				t.Error("In() should represent the same instant in time")
			}
		})
	}
}

func TestFormatTime(t *testing.T) {
	// Fixed time: January 15, 2026 at 12:00:00 UTC (winter, no DST)
	winter := time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)
	// Fixed time: July 15, 2026 at 12:00:00 UTC (summer, DST active)
	summer := time.Date(2026, 7, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		zone string
		time time.Time
		want string
	}{
		{
			name: "Denver in winter",
			zone: "America/Denver",
			time: winter,
			want: "Thu Jan 15 2026 05:00:00 GMT-0700 (Mountain Standard Time)",
		},
		{
			name: "Denver in summer",
			zone: "America/Denver",
			time: summer,
			want: "Wed Jul 15 2026 06:00:00 GMT-0600 (Mountain Daylight Time)",
		},
		{
			name: "Tokyo",
			zone: "Asia/Tokyo",
			time: winter,
			want: "Thu Jan 15 2026 21:00:00 GMT+0900 (Japan Standard Time)",
		},
		{
			name: "New York in winter",
			zone: "America/New_York",
			time: winter,
			want: "Thu Jan 15 2026 07:00:00 GMT-0500 (Eastern Standard Time)",
		},
		{
			name: "New York in summer",
			zone: "America/New_York",
			time: summer,
			want: "Wed Jul 15 2026 08:00:00 GMT-0400 (Eastern Daylight Time)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz, _ := GetTimezone(tt.zone)
			got := tz.FormatTime(tt.time)

			if got != tt.want {
				t.Errorf("FormatTime() = %q, want %q", got, tt.want)
			}
		})
	}
}
