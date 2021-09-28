package models

type Package struct {
	PackageName  string `json:"package-name"`
	Version      string `json:"version"`
	Author       string `json:"author"`
	License      string `json:"license"`
	Instructions string `json:"instructions"`
	GitHubUrl    string `json:"github-url"`
	AuthID       string `json:"authid"`
}
