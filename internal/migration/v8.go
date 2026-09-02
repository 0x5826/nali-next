package migration

import (
	"log"
	"strings"

	"github.com/spf13/viper"
	"github.com/zu1k/nali/internal/constant"
	"github.com/zu1k/nali/internal/db"
	"github.com/zu1k/nali/pkg/cdn"
	"github.com/zu1k/nali/pkg/geoip"
	"github.com/zu1k/nali/pkg/ip2region"
	"github.com/zu1k/nali/pkg/qqwry"
)

func migration2v8() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(constant.ConfigDirPath)

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	dbList := db.List{}
	err = viper.UnmarshalKey("databases", &dbList)
	if err != nil {
		log.Fatalln("Config invalid:", err)
	}

	needOverwrite := false
	for _, adb := range dbList {
		switch adb.Name {
		case "qqwry":
			if len(adb.DownloadUrls) == 0 {
				needOverwrite = true
				adb.DownloadUrls = qqwry.DownloadUrls
			} else {
				for _, u := range adb.DownloadUrls {
					if strings.Contains(u, "HMBSbige") ||
						strings.Contains(u, "FW27623") ||
						strings.Contains(u, "zu1k.com") ||
						strings.Contains(u, "99wry.cf") ||
						strings.Contains(u, "sspanel-uim") {
						needOverwrite = true
						adb.DownloadUrls = qqwry.DownloadUrls
						break
					}
				}
			}
		case "geoip":
			if len(adb.DownloadUrls) == 0 {
				needOverwrite = true
				adb.DownloadUrls = geoip.DownloadUrls
			}
		case "cdn":
			if len(adb.DownloadUrls) == 0 {
				needOverwrite = true
				adb.DownloadUrls = cdn.DownloadUrls
			}
		case "ip2region":
			if len(adb.DownloadUrls) == 0 {
				needOverwrite = true
				adb.DownloadUrls = ip2region.DownloadUrls
			}
		}
	}

	if needOverwrite {
		viper.Set("databases", dbList)
		err = viper.WriteConfig()
		if err != nil {
			log.Println(err)
		}
	}
}
