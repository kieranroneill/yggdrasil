package config

import "github.com/kirsle/configdir"

// DefaultConfigPath Gets the default config directory path.
//
// Returns:
//   - The default user config path.
func DefaultConfigPath() (string, error) {
	configPath := configdir.LocalConfig("yggdrasil")
	err := configdir.MakePath(configPath) // ensure it exists.
	if err != nil {
		return "", err
	}

	return configPath, nil
}
