package migration

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/zu1k/nali/internal/constant"
	"github.com/zu1k/nali/internal/db"
)

func TestMigration2v8_OldConfigUpgrade(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "nali-test-migration-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	constant.ConfigDirPath = tempDir
	viper.Reset()

	oldConfigContent := `
databases:
  - name: qqwry
    format: qqwry
    file: qqwry.dat
    download-urls:
      - https://gh-release.zu1k.com/HMBSbige/qqwry/qqwry.dat
  - name: geoip
    format: mmdb
    file: GeoLite2-City.mmdb
`
	err = os.WriteFile(filepath.Join(tempDir, "config.yaml"), []byte(oldConfigContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	migration2v8()

	viper.Reset()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(tempDir)
	err = viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}

	var dbs db.List
	err = viper.UnmarshalKey("databases", &dbs)
	if err != nil {
		t.Fatal(err)
	}

	for _, d := range dbs {
		if d.Name == "qqwry" {
			if len(d.DownloadUrls) == 0 || !strings.Contains(d.DownloadUrls[0], "metowolf") {
				t.Errorf("qqwry download-urls not updated, got %v", d.DownloadUrls)
			}
		}
		if d.Name == "geoip" {
			if len(d.DownloadUrls) == 0 || !strings.Contains(d.DownloadUrls[0], "GeoLite2-City.mmdb") {
				t.Errorf("geoip download-urls not updated, got %v", d.DownloadUrls)
			}
		}
	}
}
