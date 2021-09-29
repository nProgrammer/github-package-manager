package controllers

import (
	"fmt"
	"ghpm/config"
	"ghpm/models"
	"ghpm/tools"
	"strings"
)

func CheckAuth(url string, data models.LicensedPackages) bool {
	fmt.Println(data)
	dataP := config.LoadPackageInfo()
	hashURL := tools.EncSHA(url)
	hashP := strings.Split(dataP.AuthID, "-")
	fmt.Println(hashP)
	fmt.Println(hashURL)
	i := 0
	for i < len(data.Packages) {
		if data.Packages[i].AuthID == dataP.AuthID && hashURL == hashP[2] {
			return true
		}
		i++
	}
	return false
}
