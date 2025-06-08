package api

type VersionsResponse struct {
	ConfigPath string `json:"configPath"`
	Version    string `json:"version"`
}
