package config

import (
	"encoding/json"
	"ghpm/models"
	"io/ioutil"
)

func LoadAuthorizedPackagesList() models.LicensedPackages {
	file, _ := ioutil.ReadFile("./authos.json")
	data := models.LicensedPackages{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
