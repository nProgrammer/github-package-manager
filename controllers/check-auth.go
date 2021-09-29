package controllers

import (
	"encoding/json"
	"fmt"
	"ghpm/config"
	"ghpm/models"
	"ghpm/tools"
	"io/ioutil"
	"strings"
)

func CheckAuth(url string) {
	file, _ := ioutil.ReadFile("./authos.json")
	data := models.LicensedPackages{}
	_ = json.Unmarshal([]byte(file), &data)
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
