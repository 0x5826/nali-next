package geoip

import (
	"strings"
	"testing"
)

func TestCheckFile(t *testing.T) {
	if CheckFile(nil) {
		t.Error("expected false for nil data")
	}
	if CheckFile([]byte("invalid mmdb content")) {
		t.Error("expected false for invalid data")
	}
}

func TestDownloadUrls(t *testing.T) {
	if len(DownloadUrls) == 0 {
		t.Fatal("DownloadUrls should not be empty")
	}
	for _, u := range DownloadUrls {
		if !strings.HasPrefix(u, "https://") {
			t.Errorf("expected HTTPS URL, got %s", u)
		}
		if !strings.HasSuffix(u, "GeoLite2-City.mmdb") {
			t.Errorf("expected URL ending with GeoLite2-City.mmdb, got %s", u)
		}
	}
}

func TestResult_String(t *testing.T) {
	r1 := Result{
		Country:     "中国",
		CountryCode: "CN",
		Area:        "广东",
	}
	if r1.String() != "中国 广东" {
		t.Errorf("unexpected string representation: %s", r1.String())
	}

	r2 := Result{
		Country:     "Australia",
		CountryCode: "AU",
		Area:        "",
	}
	if r2.String() != "Australia" {
		t.Errorf("unexpected string representation: %s", r2.String())
	}
}
