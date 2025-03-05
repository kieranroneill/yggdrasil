package types

type AppMetadata struct {
	Environment string `yaml:"environment"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Version     string `json:"version"`
}
