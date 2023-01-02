package envtools

import (
	"github.com/joho/godotenv"
	"github.com/luaxlou/goutils/tools/envutils"
	"github.com/luaxlou/goutils/tools/fileutils"
	"log"
)

func Init() {

	if envutils.IsDev() {

		checkPaths := []string{
			"./.env",
			"../.env",
			"../../.env",
			"../../../.env",
			"../../../../.env",
			"../../../../../.env",
		}

		for _, p := range checkPaths {
			if fileutils.Exists(p) {
				err := godotenv.Load(p)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
		}

	}
}
