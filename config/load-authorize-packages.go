package config

import (
	"encoding/json"
	"ghpm/models"
	"io/ioutil"
	"os/exec"
)

func LoadAuthorizedPackagesList() models.LicensedPackages {
	cmd := exec.Command("wget", "https://raw.githubusercontent.com/nProgrammer/github-package-manager/main/authorized-packages.json")
	cmd.Run()
	file, _ := ioutil.ReadFile("./authorized-packages.json")
	data := models.LicensedPackages{}
	_ = json.Unmarshal([]byte(file), &data)
	cmd = exec.Command("rm", "./authorized-packages.json")
	cmd.Run()

	return data
}
