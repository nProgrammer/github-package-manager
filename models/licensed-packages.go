package models

type LicensedPackage struct {
	AuthID     string `json:"authid"`
	ProjectUrl string `json:"project-url"`
}

type LicensedPackages struct {
	Packages []LicensedPackage `json:"packages"`
}
