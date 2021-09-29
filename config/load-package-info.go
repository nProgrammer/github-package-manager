package config

import (
	"encoding/json"
	"ghpm/models"
	"io/ioutil"
)

func LoadPackageInfo() models.Package {
	file, _ := ioutil.ReadFile("./package-info.json")
	data := models.Package{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
