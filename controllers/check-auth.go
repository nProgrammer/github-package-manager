package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"ghpm/config"
	"ghpm/models"
	"io/ioutil"
	"strings"
)

func CheckAuth(url string) {
	file, _ := ioutil.ReadFile("./authos.json")
	data := models.LicensedPackages{}
	_ = json.Unmarshal([]byte(file), &data)
	fmt.Println(data)
	dataP := config.LoadPackageInfo()
	hash := sha256.Sum256([]byte(url))
	hashURL := hex.EncodeToString(hash[:])
	hashP := strings.Split(dataP.AuthID, "-")
	fmt.Println(hashP)
	fmt.Println(hashURL)
	if data.Packages[0].AuthID == dataP.AuthID && hashURL == hashP[2] {
		fmt.Println("Authorized")
	} else {
		fmt.Println("NON")
	}
}
