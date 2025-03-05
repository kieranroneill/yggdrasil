package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kieranroneill/yggdrasil/libs/constants"
	"github.com/kieranroneill/yggdrasil/libs/types"
	"log/slog"
	"net/http"
)

// RegisterApp Registers an app with the registry service at <registryBaseURL>/apps/register.
//
// Parameters:
//   - registryBaseURL: The base URL for the registry service.
//   - metadata: The app's metadata.
//
// Returns:
//   - An error if something went wrong.
func RegisterApp(registryBaseURL string, metadata types.AppMetadata) error {
	body, err := json.Marshal(metadata)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to register app %s: %v", metadata.Name, err))
		return err
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", registryBaseURL, constants.AppsPath, constants.RegisterPath), bytes.NewReader(body))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to register app %s: %v", metadata.Name, err))

		return err
	}

	// post the request
	_, err = http.DefaultClient.Do(request)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to register app %s: %v", metadata.Name, err))

		return err
	}

	slog.Debug(fmt.Sprintf("successfully registered app %s", metadata.Name))

	return nil
}
