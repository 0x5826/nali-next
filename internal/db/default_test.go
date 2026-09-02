package db

import (
	"testing"
)

func TestGetDefaultDBList_GeoIP(t *testing.T) {
	dbList := GetDefaultDBList()
	var geoipDB *DB
	for _, item := range dbList {
		if item.Name == "geoip" {
			geoipDB = item
			break
		}
	}

	if geoipDB == nil {
		t.Fatal("geoip database not found in default list")
	}

	if len(geoipDB.DownloadUrls) == 0 {
		t.Error("geoip DownloadUrls should not be empty")
	}

	if geoipDB.Format != FormatMMDB {
		t.Errorf("expected format %s, got %s", FormatMMDB, geoipDB.Format)
	}
}

func TestDbNameListForUpdate_GeoIP(t *testing.T) {
	found := false
	for _, name := range DbNameListForUpdate {
		if name == "geoip" {
			found = true
			break
		}
	}
	if !found {
		t.Error("geoip should be in DbNameListForUpdate")
	}

	if _, ok := DbCheckFunc[FormatMMDB]; !ok {
		t.Error("DbCheckFunc should have entry for FormatMMDB")
	}
}
