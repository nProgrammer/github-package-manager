package config

import (
	"encoding/json"
	"ghpm/models"
	"io/ioutil"
	"os"
	"os/exec"
)

func LoadAuthorizedPackagesList() models.LicensedPackages {
	cmd := exec.Command("wget", "https://github.com/nProgrammer/github-package-manager/blob/main/authorized-packages.json")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
	file, _ := ioutil.ReadFile("./authorized-packages.json")
	data := models.LicensedPackages{}
	_ = json.Unmarshal([]byte(file), &data)
	cmd = exec.Command("rm", "./authorized-packages.json")
	cmd.Run()

	return data
}
