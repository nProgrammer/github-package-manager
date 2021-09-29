package controllers

import (
	"fmt"
	"ghpm/config"
	"ghpm/models"
	"ghpm/tools"
	"strings"
)

func CheckAuth(url string, data models.LicensedPackages) {
	fmt.Println(data)
	dataP := config.LoadPackageInfo()
	hashURL := tools.EncSHA(url)
	hashP := strings.Split(dataP.AuthID, "-")
	fmt.Println(hashP)
	fmt.Println(hashURL)
	if data.Packages[0].AuthID == dataP.AuthID && hashURL == hashP[2] {
		fmt.Println("Authorized")
	} else {
		fmt.Println("NON")
	}
}
