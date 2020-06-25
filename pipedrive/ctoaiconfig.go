package pipedrive

import ctoai "github.com/cto-ai/sdk-go"

func getPipeDriveSecret(client *ctoai.Client) string {
	pipedriveKey, err := client.Sdk.GetSecret("pipedrive_api_key")
	if err != nil {
		panic(err)
	}
	return pipedriveKey
}

// Track Program Initiation
func trackInit(client *ctoai.Client) {
	initTags := []string{"pipedrive-demo", "init"}
	metadata := map[string]interface{}{"language": "golang"}
	if err := client.Sdk.Track(initTags, "init", metadata); err != nil {
		panic(err)
	}
}

// Track API Key Retrieval
func trackTestAPIKey(client *ctoai.Client) {
	initTags := []string{"pipedrive-demo", "test-api-key"}
	metadata := map[string]interface{}{"language": "golang"}
	if err := client.Sdk.Track(initTags, "test-api-key", metadata); err != nil {
		panic(err)
	}
}

func trackFailedApiKeyTest(client *ctoai.Client) {
	initTags := []string{"pipedrive-demo", "FAILED-api-key-test"}
	metadata := map[string]interface{}{"language": "golang"}
	if err := client.Sdk.Track(initTags, "FAILED-api-key-test", metadata); err != nil {
		panic(err)
	}
}

func trackSuccessfulApiKeyTest(client *ctoai.Client) {
	initTags := []string{"pipedrive-demo", "success-api-key-test"}
	metadata := map[string]interface{}{"language": "golang"}
	if err := client.Sdk.Track(initTags, "success-api-key-test", metadata); err != nil {
		panic(err)
	}
}

func trackGetPipedriveNotes(client *ctoai.Client) {
	initTags := []string{"pipedrive-demo", "get-notes"}
	metadata := map[string]interface{}{"language": "golang"}
	if err := client.Sdk.Track(initTags, "get-notes", metadata); err != nil {
		panic(err)
	}
}
